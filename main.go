package main

import (
	"log"
	"net/http"
	"oauth-client/service"
)

func main() {
	server := service.CreateService()

	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
