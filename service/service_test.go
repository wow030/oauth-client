package service_test

import (
	"net/http"
	"net/http/httptest"
	"oauth-client/service"
	"testing"
)

func TestGuestHandler(t *testing.T) {
	service := service.CreateService()
	request := httptest.NewRequest(http.MethodGet, "/guest", nil)
	response := httptest.NewRecorder()

	service.Mux.ServeHTTP(response, request)

	want := "123123"

	got := response.Body.String()

	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
