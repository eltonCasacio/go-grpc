# gRPC – An RPC library and framework

![Badge](https://img.shields.io/static/v1?label=go&message=1.20&color=blue&style=for-the-badge&logo=Go)
![Badge](https://img.shields.io/static/v1?label=RPC&message=%20&color=blue&style=for-the-badge&logo=Google)

gRPC is a modern, open source, high-performance remote procedure call (RPC)
framework that can run anywhere. gRPC enables client and server applications to
communicate transparently, and simplifies the building of connected systems.

<table>
  <tr>
    <td><b>Homepage:</b></td>
    <td><a href="https://grpc.io/">grpc.io</a></td>
  </tr>
  <tr>
    <td><b>Mailing List:</b></td>
    <td><a href="https://groups.google.com/forum/#!forum/grpc-io">grpc-io@googlegroups.com</a></td>
  </tr>
</table>

## Concepts

See [gRPC Concepts](CONCEPTS.md)

## About this repository

This is a simple CRUD to register courses and categories through the gRPC communication.

## Running project
### Prerequisites
- Go, any one of the three latest major [releases of Go](https://golang.org/doc/devel/release.html).
  
  For installation instructions, see Go’s [Getting Started](https://golang.org/doc/install) guide.

- [Protocol buffer](https://developers.google.com/protocol-buffers) compiler, protoc, [version 3](https://protobuf.dev/programming-guides/proto3).

  For installation instructions, see Protocol Buffer Compiler Installation.

- Go plugins for the protocol compiler:

1. Install the protocol compiler plugins for Go using the following commands:

    > `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
    > 
    > `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

2. Update your PATH so that the protoc compiler can find the plugins:

    > `export PATH="$PATH:$(go env GOPATH)/bin"`

- Evans

  For simulation client.

  see installation: https://github.com/ktr0731/evans#docker-image

- Sqlite
  For database

### So, let's go

- clone this repository
- into project run `go mod tidy`
- create database and tables
  - `sqlite3 db.sqlite`
  - `create table categories (id string, name string, description string);`
  - `create table courses (id string, name string, description string, category_id string);`
- start go server, `go run cmd/server/main.go`
- start evans by running: `evans -r repl`



