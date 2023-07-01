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

## Run this project
To execute this project we'll use Evans and Sqlite <br />

if you dont't have Evens installed, see: https://github.com/ktr0731/evans#docker-image

So, let's start

- clone this repository
- into project run `go mod tidy`
- create database and tables
  - `sqlite3 db.sqlite`
  - `create table categories (id string, name string, description string);`
  - `create table courses (id string, name string, description string, category_id string);`
- start go server, `go run cmd/server/main.go`
- start evans by running: `evans -r repl`



