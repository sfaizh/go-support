# Support ticket system written in Go

Go Support is an API service that handles ticket management.

## Features

- Modular design with structured packages
- Built-in logging and filesystem tools
- Ticket management CRUD operations
- Database interaction (Postgres) via internal helpers

## Requirements

- Go 1.18 or later

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/go-support.git
   cd go-support-main
   ```

2. **Build the CLI binary:**
   ```bash
   make build
   ```

   Or manually using:
   ```bash
   go build -o bin/gosupport ./cmd/core
   ```

3. **Run the CLI:**
   ```bash
   ./bin/gosupport
   ```

## Project Structure

```
go-support-main/
├── cmd/
│   ├── cli/              # CLI entry point
│   └── core/             # Core logic entry point
├── internal/
│   ├── globals/          # Global configurations
│   ├── structs/          # Struct definitions and defaults
│   ├── ticket/           # Ticket logic
│   └── util/             # Utility packages
│       ├── api/          # API helpers
│       ├── database/     # Database operations
│       ├── filesystem/   # File operations
│       ├── logger/       # Logging utilities
│       └── tcpserver/    # TCP server functions
├── go.mod                # Go module file
├── Makefile              # Build instructions
└── README.md             # Project info
```
