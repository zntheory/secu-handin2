package network

import (
	"crypto/tls"
	"fmt"
	"github.com/zntheory/secu-handin2/pkg/config"
	"log"
	"net"
	"sync"
)

var swg sync.WaitGroup

func sendReceiveSecrets(conn net.Conn, secret []int) (received []int) {
	// Send secret to server via conn :D
	_, errSend := fmt.Fprint(conn, secret)
	if errSend != nil {
		log.Fatalf("failed to send secrets: %v", errSend)
	}
	
}

func configureClientTLS() (tlsConfig *tls.Config, err error) {
	log.Printf("Configuring the client's TLS config.\n")

	pathCertFile := config.CertFile

	tlsConfig, err = ConfigureClientTLS(pathCertFile)
	if err != nil {
		return nil, fmt.Errorf("error configuring client TLS: %v", err)
	}
	return tlsConfig, nil
}

func ConnectToServer(secret []int) (received []int) {
	log.Printf("Client wants to connect to server.\n")

	tlsConfig, err := configureClientTLS()
	if err != nil {
		log.Fatalf("error configuring client TLS: %v", err)
	}

	conn, err := CreateConnection(tlsConfig)
	if err != nil {
		log.Fatalf("error creating connection: %v", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error closing connection: %v", err)
		}
	}(conn)
	log.Printf("Connected to server.")

	received = sendReceiveSecrets(conn, secret)
	return received
}
