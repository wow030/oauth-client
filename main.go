package main

import (
	"log"
	"net/http"
	"oauth-client/service"
)

func main() {
	server := service.CreateService()

	err := http.ListenAndServe(":80", server.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
