package config

import "path/filepath"

var (
	CertFile = filepath.Join("config", "hospital.crt")
	KeyFile  = filepath.Join("config", "hospital.key")
)
