package config

import "path/filepath"

var (
	ConnCount = 4 // Ensure only 1 hospital + 3 patients (A, B C) can join this session :)
	Port      = ":8080"
	CertFile  = filepath.Join("config", "hospital.crt")
	KeyFile   = filepath.Join("config", "hospital.key")
)
