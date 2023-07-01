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

This is a simple server that allows register courses and categories. (CRUD)

## Run this project
To execute this project some dependencies is necessary, Evans and Sqlite <br />

if you dont't hava Evens installed, you can just run <code>go install github.com/ktr0731/evans@latest</code><br />
See: https://github.com/ktr0731/evans#docker-image

So, let's start
<ol>
  <li>clone this repository</li>
  <li>enter into project and run <code>go mod tidy</code></li>
  <li>create database and tables...</li>
  <li>start go server, <code>go run cmd/server/main.go</code></li>
  <li>start evans by running: <code>evans -r repl</code></li>
</ol>


## Some used commands

create struct gRPC
- protoc --go_out=. --go-grpc_out=.  proto/course_category.proto


evans
- https://github.com/ktr0731/evans#macos

- If your server is enabling gRPC reflection, you can launch Evans with only -r (--reflection) option.
evans -r repl
call


sqlite
- sqlite3 db.sqlite
- create table categories (id string, name string, description string);





