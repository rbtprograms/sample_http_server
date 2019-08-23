package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testGETPlayers(t *testing.T) {
	t.Run("returns Peppers score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}