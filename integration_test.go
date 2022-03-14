package main

import (
	"fmt"
	"golang-tdd/http_server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetreive(t *testing.T) {
	store := http_server.NewInMemoryPlayerStore()
	server := http_server.NewPlayerServer(store)
	player := "Pepper"

	request, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", player), nil)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)

	response := httptest.NewRecorder()
	request, _ = http.NewRequest("GET", fmt.Sprintf("/players/%s", player), nil)
	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("got %d want %d", response.Code, 200)
	}
	if response.Body.String() != "3" {
		t.Errorf("got %s want %s", response.Body.String(), "3")
	}
}
