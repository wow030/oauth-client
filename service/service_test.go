package service_test

import (
	"net/http"
	"net/http/httptest"
	"oauth-client/service"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("TOKEN_URL", "http://localhost:8001/token")
	os.Exit(m.Run())
}
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
