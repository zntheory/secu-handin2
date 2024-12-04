# secu-handin2

## Quickstart guide

WIP.

## Structure overview
```txt
secu-handin2/
├── cmd/
│   ├── hospital/         # Hospital server application
│   └── patient/          # Patient client application
├── docs/
│   ├── report.pdf         # Touches on GDPR, adversarial model, solution design and guarantees.
│   └── usage.md           # How to run the application
├── pkg/
│   ├──config/
│   │   └── config.go      # Application configuration (e.g., ports, participant settings)
│   ├──network/
│   │   ├── certauth.go    # Certification Authority logic
│   │   ├── client.go      # Network client implementation for patients
│   │   ├── server.go      # Network server implementation for the hospital
│   │   └── tls.go         # TLS setup and key exchange
│   └──sharing/
│       ├── aggregation.go # Secure aggregation logic
│       ├── shares.go      # Additive secret sharing implementation
│       └── utils.go       # Helper functions for sharing (e.g., modular arithmetic)
├── go.mod                 # Go module dependencies
└── README.md              # Project overview and quickstart guide
```