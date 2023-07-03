FROM golang:1.20
WORKDIR /app
COPY . .

RUN apt update && \
    apt install sqlite3 && \
    apt install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install github.com/ktr0731/evans@latest && \
    go mod tidy

CMD [ "tail", "-f", "/dev/null" ]

# must be included export PATH=$PATH:$(go env GOPATH)/bin in .profile or run in terminal into container