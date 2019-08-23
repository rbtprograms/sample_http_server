package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testGETPlayers(t *testing.T) {
	server := &PlayerServer{}
	t.Run("returns Peppers score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})
	
	t.Run("returns Floyds score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response:= httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	route := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodGet, route, nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}