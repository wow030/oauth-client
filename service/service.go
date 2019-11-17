package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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

	return &service
}

func (s *Service) GuestHandler(w http.ResponseWriter, r *http.Request) {
	// redirect to auth-server
	// use post

	client := &http.Client{}

	values := make(url.Values)

	values.Add("grant_type", "client_credentials")

	resp, err := client.PostForm(s.tokenURL, values)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	w.Write([]byte("123123"))

}
