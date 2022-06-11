package main

import (
	"net"
	"os"

	"github.com/kanguki/grpc-microservices-example/div/div"
	"github.com/kanguki/grpc-microservices-example/div/log"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("DIV_PORT")
	if port == "" {
		port = ":4003"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Log("error listening to port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	divService := div.Service{Max: 1000, Min: -1000, MaxFreeTier: 10, MinFreeTier: -10}
	div.RegisterDivServer(grpcServer, divService)
	log.Log("Starting grpc service on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("error serving grpc server: %v", err)
	}
}
