package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// permet de creer une nouvelle requete
		// le premier argument est la requete et le second le path
		// le nil refer to le request body
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// ajout d'un spy
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
