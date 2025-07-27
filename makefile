deps:
	go mod tidy
	go mod vendor

rpc:
	protoc --go_out=pb --go-grpc_out=pb proto/book.proto

lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o bookservice cmd/main.go

clean:
	go clean

.PHONY: deps rpc lint test build clean
.DEFAULT_GOAL := build


