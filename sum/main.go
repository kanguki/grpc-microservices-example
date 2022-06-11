package main

import (
	"net"
	"os"

	"github.com/kanguki/grpc-microservices-example/sum/log"
	"github.com/kanguki/grpc-microservices-example/sum/sum"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("SUM_PORT")
	if port == "" {
		port = ":4000"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Log("error listening to port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	sumService := sum.Service{Max: 1000, Min: -1000, MaxFreeTier: 10, MinFreeTier: -10}
	sum.RegisterSumServer(grpcServer, sumService)
	log.Log("Starting grpc service on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("error serving grpc server: %v", err)
	}
}
