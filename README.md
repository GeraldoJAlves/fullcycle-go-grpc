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

## Test with evans
 - https://github.com/ktr0731/evans?tab=readme-ov-file#installation
 - `evans -r repl`
 - `package pb` selecting package
 - `service CategoryService` using CategoryService
 - `call ListCategories` list all categories

## Proto grpc
 - https://grpc.io/docs/languages/go/quickstart/
