package network

import (
	"crypto/tls"
	"fmt"
	"github.com/zntheory/secu-handin2/pkg/config"
	"log"
	"math/rand"
	"net"
	"sync"
)

// Client holds the ID, Conn and Status of a patient.go instance
type Client struct {
	ID     string
	Conn   net.Conn
	Status int
	// 0: Connected.
	// 1: Sent secret shares.
	// 2: Sent secrets and received config.ConnCount secrets.
}

// clients is a map of connected patient.go instances
var clientsByID = make(map[string]*Client)
var clientsByConn = make(map[net.Conn]*Client)
var mu sync.Mutex
var wg sync.WaitGroup
var cg sync.WaitGroup
var clientCount = 0

// generateClientID and Client.ID could be deleted -> conn as ID...
func generateClientID() (clientID string) {
	for {
		clientID = fmt.Sprintf("c-%d", rand.Intn(100000)) // Random ClientID
		if _, exists := clientsByID[clientID]; !exists {
			break
		}
	}
	return clientID
}

// createClient using a clientID and conn
func createClient(clientID string, conn net.Conn) (client *Client) {
	client = &Client{
		ID:     clientID,
		Conn:   conn,
		Status: 0,
	}
	return client
}

func addClient(conn net.Conn) {
	mu.Lock()
	if clientCount < config.ConnCount {
		client := createClient(generateClientID(), conn)
		clientsByID[client.ID] = client
		clientsByConn[client.Conn] = client
		clientCount++

		log.Printf("Added client %s. Patient count: %d.", client.ID, clientCount)
		fmt.Fprintf(conn, "Welcome to the SECU Hospital server!\n"+
			"Waiting for all clients to connect before receiving secret shares.\n"+
			"Current no. of clients: %d", clientCount)
		wg.Done()
	} else {
		fmt.Fprint(conn, "Room full. Closing connection.\n- Please contact the SECU hospital if this is a mistake.")
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}
	mu.Unlock()
}

func secretShare(conn net.Conn) {
	var secret int
	_, err := fmt.Fscanf(conn, "%d", &secret)
	if err != nil {
		log.Printf("Error reading secret share: %v", err)
	}

	mu.Lock()
	for _, client := range clientsByConn {
		if client.Conn != conn {
			fmt.Fprintf(client.Conn, "%d", secret)
			log.Printf("Secret %d shared to client %s.\n", secret, client.ID)
		}
	}
	mu.Unlock()
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		log.Printf("Closing connection for client %s via defer", clientsByConn[conn].ID)
		err := conn.Close()
		if err != nil {
			log.Fatalf("error closing connection: %v", err)
		}
	}(conn)
	addClient(conn)
	wg.Wait()
	secretShare(conn)
}

// configureServerTLS uses the hospital's cert and pk
func configureServerTLS() (tlsConfig *tls.Config, err error) {
	log.Printf("Configuring the server's TLS config.\n")

	tlsConfig, err = ConfigureServerTLS(config.CertFile, config.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("error configuring server TLS: %v", err)
	}
	return tlsConfig, nil
}

func StartServer() {
	log.Printf("Starting server to fit %d clients.\n", config.ConnCount)
	wg.Add(config.ConnCount)

	tlsConfig, err := configureServerTLS()
	if err != nil {
		log.Fatalf("error configuring TLS for server: %v", err)
	}

	listener, err := CreateListener(tlsConfig)
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("error closing listener: %v", err)
		}
	}(listener)

	log.Printf("Server is listening on %v with TLS", config.Port)
	for {
		conn, errListen := listener.Accept()
		if errListen != nil {
			log.Printf("Error accepting connection: %v\n", errListen)
			continue
		}
		go handleConnection(conn)
	}
}
