# secu-handin2

## Quickstart guide

WIP.

## Structure overview
```
secu-handin2/
├── cmd/
│   ├── hospital/         # Hospital server application
│   └── patient/          # Patient client application
├── pkg/
│   ├── sharing/
│   │   ├── shares.go      # Additive secret sharing implementation
│   │   ├── aggregation.go # Secure aggregation logic
│   │   └── utils.go       # Helper functions for sharing (e.g., modular arithmetic)
│   ├── network/
│   │   ├── server.go      # Network server implementation for the hospital
│   │   ├── client.go      # Network client implementation for patients
│   │   └── tls.go         # TLS setup and key exchange
│   └── config/
│       └── config.go      # Application configuration (e.g., ports, participant settings)
├── docs/
│   ├── report.pdf         # Touches on GDPR, adversarial model, solution design and guarantees.
│   └── usage.md           # How to run the application
├── go.mod                 # Go module dependencies
└── README.md              # Project overview and quickstart guide
```