# What

Service that can mul 2 numbers.

# Usage

go run main.go

# Dev note

gen 1 file. I used this

```sh
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative mul/mul.proto
```

but if you want to gen 2 files :D

```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mul/mul.proto
```
