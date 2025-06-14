package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestGETPlayers(t *testing.T) {
//	t.Run("returns Pepper's score", func(t *testing.T) {
//		// permet de creer une nouvelle requete
//		// le premier argument est la requete et le second le path
//		// le nil refer to le request body
//		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
//		// ajout d'un spy
//		response := httptest.NewRecorder()
//
//		PlayerServer(response, request)
//
//		got := response.Body.String()
//		want := "20"
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//
//	t.Run("returns Floyd's score", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
//		response := httptest.NewRecorder()
//
//		PlayerServer(response, request)
//
//		got := response.Body.String()
//		want := "10"
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//}

//func TestGetPlayerScore(t *testing.T) {
//	server := &PlayerServer{}
//
//	t.Run("returns Pepper's score", func(t *testing.T) {
//		// fonction qui permet de faire la requete
//		request := newGetScoreRequest("Pepper")
//		// permet de mocker la reponse HTTP
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		assertResponseBody(t, response.Body.String(), "20")
//	})
//
//	t.Run("returns Floyd's score", func(t *testing.T) {
//		request := newGetScoreRequest("Floyd")
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		assertResponseBody(t, response.Body.String(), "10")
//	})
//}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

//type StubPlayerStore struct {
//	scores map[string]int
//}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

//func TestGETPlayers(t *testing.T) {
//	store := StubPlayerStore{
//		map[string]int{
//			"Pepper": 20,
//			"Floyd":  10,
//		},
//	}
//	server := &PlayerServer{&store}
//
//	t.Run("returns Pepper's score", func(t *testing.T) {
//		request := newGetScoreRequest("Pepper")
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		assertResponseBody(t, response.Body.String(), "20")
//	})
//
//	t.Run("returns Floyd's score", func(t *testing.T) {
//		request := newGetScoreRequest("Floyd")
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		assertResponseBody(t, response.Body.String(), "10")
//	})
//
//	t.Run("return 404 on missing players", func(t *testing.T) {
//		// on creer une nouvelle requete avec un player qui n'existe pas
//		request := newGetScoreRequest("Apollo")
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		// on definis les status code que l'on doit recuperer si le player existe pas
//		got := response.Code
//		want := http.StatusNotFound
//
//		if got != want {
//			t.Errorf("got status %d want %d", got, want)
//		}
//	})
//}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		// on creer une nouvelle requete avec un player qui n'existe pas
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", want, got)
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it records wins when POST", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

//func TestStoreWins(t *testing.T) {
//	store := StubPlayerStore{
//		map[string]int{},
//	}
//	server := &PlayerServer{&store}
//
//	t.Run("it returns accepted on POST", func(t *testing.T) {
//		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		assertStatus(t, response.Code, http.StatusAccepted)
//	})
//}
