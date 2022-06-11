package main

import (
	"net"
	"os"

	"github.com/kanguki/grpc-microservices-example/mul/log"
	"github.com/kanguki/grpc-microservices-example/mul/mul"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("MUL_PORT")
	if port == "" {
		port = ":4002"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Log("error listening to port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	mulService := mul.Service{}
	mul.RegisterMulServer(grpcServer, mulService)
	log.Log("Starting grpc service on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("error serving grpc server: %v", err)
	}
}
