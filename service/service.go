package service

import (
	"net/http"
)

type Service struct {
	Mux *http.ServeMux
}

func CreateService() *Service {
	service := Service{}

	service.Mux = http.NewServeMux()

	service.Mux.HandleFunc("/guest", GuestHandler)

	return &service
}

func GuestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("123123"))

}
