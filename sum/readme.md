# What

Service that can sum 2 numbers.

# Usage

go run main.go

# Dev note

PROTOC

gen 1 file. I used this

```sh
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative sum/sum.proto
```

but if you want to gen 2 files :D

```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sum/sum.proto
```

CONTAINERIZING

```sh
docker build . -t grpc-microservice-example-sum:1.0.0
```

```sh
docker run -it --rm --name=sum-ex grpc-microservice-example-sum:latest
```
