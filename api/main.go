package main

import (
	"github.com/kanguki/grpc-microservices-example/api/apiServer"
	"github.com/kanguki/grpc-microservices-example/api/auth"
	"github.com/kanguki/grpc-microservices-example/api/log"
	"net/http"
)

func main() {
	apiServer, err := apiServer.NewServer()
	if err != nil {
		log.Log("error starting api server: %v", err)
	}

	http.HandleFunc("/", apiServer.Auth(func(writer http.ResponseWriter, request *http.Request) {
	}, []auth.ACL{auth.ACL_READ}))

	log.Log("Starting server on port %v\n", apiServer.Port)
	if err := http.ListenAndServe(apiServer.Port, nil); err != nil {
		log.Log("error serving api server: %v", err)
	}
}
