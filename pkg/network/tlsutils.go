package network

import (
	"crypto/tls"
	"fmt"
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
	cert, err := LoadCertificate(certFile)
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
