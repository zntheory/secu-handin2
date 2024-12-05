package network

import (
	"crypto/tls"
	"fmt"
	"github.com/zntheory/secu-handin2/pkg/config"
	"net"
)

// LoadCertificate for the client
func LoadCertificate(certFile string) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certFile, certFile)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to load certificate: %v", err)
	}
	return cert, nil
}

// LoadCertPK for the server
func LoadCertPK(certFile string, pkFile string) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certFile, pkFile)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to load certificate: %v", err)
	}
	return cert, nil
}

// ConfigureServerTLS for the hospital
func ConfigureServerTLS(certFile string, pkFile string) (*tls.Config, error) {
	cert, err := LoadCertPK(certFile, pkFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate during sever config: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13,
	}
	return tlsConfig, nil
}

// ConfigureClientTLS for the patient to verify the hospital's cert
func ConfigureClientTLS(certFile string) (*tls.Config, error) {
	cert, err := LoadCertificate(certFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate during client config: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
	}
	return tlsConfig, nil
}

// CreateListener uses tls.Listen and a tls.Config to create a net.Listener
func CreateListener(tlsConfig *tls.Config) (net.Listener, error) {
	listener, err := tls.Listen("https", config.Port, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %v", err)
	}
	return listener, nil
}

// CreateConnection uses tls.Dial and a tls.Config to create a net.Conn
func CreateConnection(tlsConfig *tls.Config) (net.Conn, error) {
	conn, err := tls.Dial("https", config.Port, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}
	return conn, nil
}
