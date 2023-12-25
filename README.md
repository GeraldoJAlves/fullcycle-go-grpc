# Fullcycle comunication with gRPC

## Pre requisites
 - go > 1.20
 - protoc
 - docker

## Running
 - `go mod tidy` go dependencies
 - `docker-compose up -d` starting mysql
 - `go run ./cmd/server.go` starting grpc server

## Generating proto files
 - `protoc --go_out=. --go-grpc_out=. ./proto/course_category.proto`

## Proto grpc
 - https://grpc.io/docs/languages/go/quickstart/
