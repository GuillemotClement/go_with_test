package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore store score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

//func PlayerServer(w http.ResponseWriter, r *http.Request) {
//	// on recupere le joueur depuis l'url
//	player := strings.TrimPrefix(r.URL.Path, "/players/")
//	// w est un objet fournis par le serveur HTTP pour ecrire la reponse HTTP
//	// c'est lui qui permet d'ecrire du text au client
//	fmt.Fprint(w, GetPlayerScore(player))
//}

//func GetPlayerScore(name string) string {
//	if name == "Pepper" {
//		return "20"
//	}
//
//	if name == "Floyd" {
//		return "10"
//	}
//
//	return ""
//}
