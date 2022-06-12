package main

import (
	"github.com/gorilla/mux"
	"github.com/kanguki/grpc-microservices-example/api/apiServer"
	"github.com/kanguki/grpc-microservices-example/api/auth"
	"github.com/kanguki/grpc-microservices-example/api/log"
	"net/http"
)

func main() {
	server, err := apiServer.NewServer()
	if err != nil {
		log.Log("error starting api server: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	r.HandleFunc("/login", server.Login).Methods("POST")
	r.HandleFunc("/logout", server.Logout).Methods("POST")
	r.HandleFunc("/{type:(?:sum|sub|mul|div)}", server.Auth(server.Calculate, []auth.ACL{auth.ACL_READ}))

	log.Log("Starting server on port %v\n", server.Port)
	if err := http.ListenAndServe(server.Port, r); err != nil {
		log.Log("error serving api server: %v", err)
	}
}
