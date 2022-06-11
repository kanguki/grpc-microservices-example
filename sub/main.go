package main

import (
	"net"
	"os"

	"github.com/kanguki/grpc-microservices-example/sub/log"
	"github.com/kanguki/grpc-microservices-example/sub/sub"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("SUB_PORT")
	if port == "" {
		port = ":4001"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Log("error listening to port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	subService := sub.Service{Max: 1000, Min: -1000, MaxFreeTier: 10, MinFreeTier: -10}
	sub.RegisterSubServer(grpcServer, subService)
	log.Log("Starting grpc service on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("error serving grpc server: %v", err)
	}
}
