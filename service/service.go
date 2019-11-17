package service

import (
	"fmt"
	"net/http"
	"os"
)

type Service struct {
	Mux      *http.ServeMux
	tokenURL string
}

type AuthConfig struct {
	TokenURL string
}

func CreateService() *Service {
	service := Service{}

	service.Mux = http.NewServeMux()

	service.Mux.HandleFunc("/guest", service.GuestHandler)

	service.tokenURL = os.Getenv("TOKEN_URL")

	fmt.Println(service.tokenURL)
	return &service
}

func (s *Service) GuestHandler(w http.ResponseWriter, r *http.Request) {
	// redirect to auth-server
	//http.Redirect(w, r)
	w.Write([]byte("123123"))

}
