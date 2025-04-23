# gRPC Microservices in Go

This repository contains a set of Go-based microservices communicating over gRPC, showcasing Protobuf-defined service contracts, server/client implementations, and best practices for structuring a Go project with `cmd/`, `pkg/`, and `internal/` packages.

---

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Project Structure](#project-structure)
5. [Building](#building)
6. [Protobuf Definitions](#protobuf-definitions)
7. [Generating Code](#generating-code)
8. [Running Services](#running-services)
9. [Client Usage](#client-usage)
10. [Configuration](#configuration)
11. [Testing](#testing)
12. [CI/CD](#cicd)
13. [Contributing](#contributing)
14. [License](#license)

---

## Overview

This project demonstrates a pattern for building and organizing Go microservices that communicate using gRPC:

- **Protobuf** (`proto/`) for service definitions and message formats.
- **Server** implementations in `cmd/`, each exposing gRPC endpoints.
- **Client** and utility code in `pkg/` for easy consumption of services.
- **Internal** packages for non-exported application logic.
- **Local** sandbox environment for running with sample data.

## Features

- gRPC services with **Gin**-style interceptors for logging and recovery.
- **TLS** and **authentication** interceptor examples.
- **Health** service for readiness and liveness checks.
- Clean project layout following [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Prerequisites

- Go 1.18+ installed
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` & `protoc-gen-go-grpc` installed (`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` and `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`)
- GNU Make (for convenience)

## Project Structure

```text
server/
├── Makefile            # Build and code-generation commands
├── build/              # CI/CD build scripts and configs
├── cmd/                # Entrypoints for each service
│   ├── service1/       # main.go for Service1
│   └── service2/       # main.go for Service2
├── internal/           # Non-exported application logic
│   ├── service1/       # handlers, business logic
│   └── service2/       # handlers, business logic
├── pkg/                # Exported library code (clients, helpers)
│   └── client/         # Client stubs and wrappers
├── proto/              # .proto definitions for all services
├── local/              # Local-only config (credentials, localhost certs)
├── go.mod
└── go.sum
```

## Building

```bash
# Clone
git clone https://github.com/nicktretyakov/server.git
cd server

# Generate Protobuf code
make proto

# Build all services
make build
```

## Protobuf Definitions

All service definitions live under `proto/` with appropriate `.proto` files. Example:

```proto
syntax = "proto3";
package example;

service Greeter {
  rpc SayHello(HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

## Generating Code

Run:

```bash
make proto
```

This invokes:

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

and outputs generated `.pb.go` files adjacent to the `.proto` sources.

## Running Services

Each service has its own `cmd/{service}` directory. For example, to run `service1`:

```bash
cd cmd/service1
go run main.go --config ../../local/service1.yaml
```

Services listen on ports defined in their YAML configs and register with the health service.

## Client Usage

The `pkg/client` package includes generated client wrappers. Example usage:

```go
import (
  "context"
  "log"
  "github.com/nicktretyakov/server/pkg/client"
)

func main() {
  conn, err := client.Dial("localhost:50051")
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  greeter := client.NewGreeterClient(conn)
  resp, err := greeter.SayHello(context.Background(), &client.HelloRequest{Name: "World"})
  if err != nil {
    log.Fatal(err)
  }
  log.Println(resp.Message)
}
```

## Configuration

Each service reads a YAML or JSON config file specifying:

```yaml
server:
  port: 50051
  tls:
    cert_file: "path/to/cert.pem"
    key_file: "path/to/key.pem"
auth:
  token: "s3cr3t"
```

Configs for local development reside under `local/`.

## Testing

```bash
# Run unit tests
go test ./internal/... ./pkg/...
```

Integration tests may be added under `test/` using `docker-compose` to spin up dependencies.

## CI/CD

The `build/` directory contains YAML for GitHub Actions or GitLab CI:

- **Lint**: `golangci-lint run`
- **Test**: `go test` with coverage
- **Build**: `go build` for each service
- **Publish**: Docker image build and push

## Contributing

Contributions are welcome. Please:

1. Fork the repo.
2. Create a feature branch: `git checkout -b feature/awesome`.
3. Make changes and add tests.
4. Submit a pull request.

Ensure code adheres to `gofmt` and `golangci-lint` standards.

## License

This project is released under the MIT License. See [LICENSE](LICENSE) for details.


