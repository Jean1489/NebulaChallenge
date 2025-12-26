# NebulaChchallenge

CLI tool for analyzing TLS/SSL security of domains using the SSL Labs API. Built with Go for the Truora Nebula Challenge.

## Description

This CLI allows you to request, monitor and retrieve SSL/TLS analysis reports for any public domain. It validates and sanitizes hostnames, interacts with the SSL Labs public API, and displays the results in a readable format or as JSON.

## Usage

Run the tool from the project root:
```bash
go run . --host=<hostname> [options]
```

### Options

- `--host string` - Hostname to analyze (required)
- `--publish` - Publish results on SSL Labs public boards
- `--json` - Output results as JSON
- `--help` - Show help message

### Examples
```bash
go run . --host=google.com
go run . --host=facebook.com --json
go run . --host=github.com --publish
```

## Features

- Hostname validation and sanitization
- Integration with SSL Labs API v2
- Polling until the analysis is completed
- Handling of rate limits and API errors
- Optional JSON output
- Graceful shutdown on Ctrl+C

## Project Architecture
```
NebulaChallenge/
│
├── main.go                 # Application entry point and CLI
├── go.mod                  # Go module definition
├── README.md               # This file
│
├── client/                 # HTTP client for SSL Labs API
│   └── ssllabs.go         # API communication logic
│
├── models/                 # Data structures
│   ├── info.go            # API info structure
│   ├── host.go            # Host analysis structure
│   ├── endpoint.go        # Endpoint structure
│   └── details.go         # Detailed endpoint information
│
├── analyzer/               # Analysis orchestration
│   └── analyzer.go        # Analysis flow and polling logic
│
├── formatter/              # Output formatting
│   └── output.go          # Text and JSON formatting
│
└── utils/                  # Helper utilities
    └── validator.go       # Input validation
```

## Requirements

- Go 1.21+
- Internet connection

## Notes

This project was developed as part of the Truora Nebula technical challenge.
