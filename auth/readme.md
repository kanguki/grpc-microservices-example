# What

Service that do authentication and authorization 😃.

# Usage

go run main.go

# Dev note

PROTOC

gen 1 file. I used this

```sh
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative auth/auth.proto
```

but if you want to gen 2 files :D

```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auth/auth.proto
```

CONTAINERIZING

```sh
docker build . -t grpc-microservice-example-auth:1.0.0
```

```sh
docker run -it --rm --name=sum-ex grpc-microservice-example-auth:latest
```
