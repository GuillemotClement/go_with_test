package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrivingThem(t *testing.T) {
	// on creer les deux composants
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(&store)
	player := "Pepper"
	// on enregistre 3 victoire pour user
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	// on fait une requete pour retrouver le nombre de victoire enregistrer
	server.ServeHTTP(response, newGetScoreRequest(player))
	// on test le status code retourner
	assertStatus(t, response.Code, http.StatusOK)
	// on test la reponse
	assertResponseBody(t, response.Body.String(), "3")
}
