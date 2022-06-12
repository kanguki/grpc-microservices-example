package main

import (
	"net"
	"os"
	"strconv"

	"github.com/kanguki/grpc-microservices-example/auth/auth"
	"github.com/kanguki/grpc-microservices-example/auth/log"
	"google.golang.org/grpc"
)

func main() {
	port, _tokenClearDuration := os.Getenv("AUTH_PORT"), os.Getenv("TOKEN_CLEAR_DURATION")
	tokenClearDuration, _ := strconv.ParseInt(_tokenClearDuration, 10, 64)
	if port == "" {
		port = ":4004"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Log("error listening to port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	db := auth.Db{Users: map[string]auth.User{
		"mo":    {Password: "mo", Acls: map[auth.ACL]bool{auth.ACL_READ: true}},
		"admin": {Password: "admin", Acls: map[auth.ACL]bool{auth.ACL_READ: true, auth.ACL_WRITE: true}},
	}}
	authService := auth.NewService(&db, tokenClearDuration)
	auth.RegisterAuthServer(grpcServer, authService)
	log.Log("Starting grpc service on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Log("error serving grpc server: %v", err)
	}
}
