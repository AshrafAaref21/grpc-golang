# gRPC Learning Project

Educational project demonstrating gRPC communication patterns in Go. Implements three services (Greet, Calculator, Blog) showcasing unary and streaming RPCs, error handling, deadlines, and persistent data operations.

**Repository**: `github.com/AshrafAaref21/grpc-golang`  
**Go Version**: 1.25.5  
**Framework**: gRPC v1.79.3 with Protocol Buffers v1.36.11

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [Services](#services)
- [Building](#building)
- [Troubleshooting](#troubleshooting)

---

## Features

| Feature | Services |
| --- | --- |
| Unary RPC | Greet, Calculator, Blog |
| Server Streaming | Greet, Calculator, Blog |
| Client Streaming | Greet, Calculator |
| Bidirectional Streaming | Greet, Calculator |
| Deadline Handling | Greet |
| Error Handling | All |
| Persistent Data (MongoDB) | Blog |
| Cross-Platform Builds | All |

---

## Prerequisites

- **Go** 1.25.5+
- **protoc** v3.21.12+ ([download](https://github.com/protocolbuffers/protobuf/releases))
- **gRPC Go plugins**:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- **Make** (build automation)
- **Docker & Docker Compose** (for Blog service only)

---

## Quick Start

```bash
git clone https://github.com/AshrafAaref21/grpc-golang.git
cd grpc-golang
go mod download
make all
```

Run any service (Terminal 1):

```bash
./bin/greet/server     # or calculator/server or blog/server
```

Run client (Terminal 2):

```bash
./bin/greet/client     # or calculator/client or blog/client
```

---

## Project Structure

```txt
grpc_course/
├── Makefile
├── go.mod
├── README.md
│
├── greet/              # Greeting service (all RPC patterns)
│   ├── proto/          # Service definitions
│   ├── server/main.go
│   └── client/main.go
│
├── calculator/         # Math operations service
│   ├── proto/          # Service definitions
│   ├── server/main.go
│   └── client/main.go
│
├── blog/               # CRUD service with MongoDB
│   ├── proto/          # Service definitions
│   ├── server/         # Server implementation
│   ├── client/main.go
│   └── docker-compose.yml
│
└── bin/                # Compiled binaries
    ├── greet/
    ├── calculator/
    └── blog/
```

---

## Services

### Greet Service (Port 50051)

Demonstrates all gRPC communication patterns:

| Method | Type | Description |
| --- | --- | --- |
| Greet | Unary | Simple greeting |
| GreetManyTimes | Server Streaming | Multiple greetings |
| LongGreet | Client Streaming | Aggregate greeting |
| GreetEveryOne | Bidirectional Streaming | Simultaneous in/out |
| GreetWithDeadline | Unary + Timeout | Deadline handling |

**Run**:

```bash
# Terminal 1
./bin/greet/server

# Terminal 2
./bin/greet/client
```

---

### Calculator Service (Port 50051)

Mathematical operations using different streaming patterns:

| Method | Type | Operation |
| --- | --- | --- |
| Sum | Unary | Add two integers |
| Primes | Server Streaming | Generate primes |
| Avg | Client Streaming | Calculate average |
| Max | Bidirectional Streaming | Track maximum |
| Sqrt | Unary | Square root |

**Run**:

```bash
# Terminal 1
./bin/calculator/server

# Terminal 2
./bin/calculator/client
```

---

### Blog Service (Port 50051)

CRUD operations with MongoDB persistence:

| Method | Type | Operation |
| --- | --- | --- |
| CreateBlog | Unary | Create post |
| ReadBlog | Unary | Retrieve post |
| UpdateBlog | Unary | Update post |
| DeleteBlog | Unary | Delete post |
| ListBlogs | Server Streaming | List all posts |

**Setup MongoDB**:

```bash
cd blog
docker-compose up -d
```

Access MongoDB UI at [http://localhost:8081](http://localhost:8081) (admin/pass)

**Run**:

```bash
# Terminal 1
./bin/blog/server

# Terminal 2
./bin/blog/client
```

---

## Building

```bash
make all              # Build all services
make greet            # Build greet service
make calculator       # Build calculator service
make blog             # Build blog service
make rebuild          # Clean and rebuild
make clean            # Remove binaries
```

---

## Troubleshooting

**Port Already in Use**:

```bash
# Linux/macOS
lsof -i :50051 | grep LISTEN | awk '{print $2}' | xargs kill

# Windows
netstat -ano | findstr :50051
taskkill /PID <PID> /F
```

**Protoc Not Found**:

```bash
# macOS (Homebrew)
brew install protobuf

# Ubuntu/Debian
sudo apt-get install protobuf-compiler

# Windows (Chocolatey)
choco install protoc
```

**gRPC Plugin Not Found**:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**MongoDB Connection Failed**:

```bash
cd blog
docker-compose up -d
docker ps | grep mongodb
```

---

## Resources

- [gRPC Documentation](https://grpc.io/docs/)
- [Protocol Buffers Guide](https://protobuf.dev/)
- [gRPC Go Examples](https://github.com/grpc/grpc-go/tree/master/examples)
- [MongoDB Documentation](https://www.mongodb.com/docs/)
- [Go Documentation](https://golang.org/doc/)
