package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore store score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// on recupere le nom du player depuis l'url
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	// on verifie la methode de la request
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// on verifie la methode de la request
//	switch r.Method {
//	case http.MethodPost:
//		p.processWin(w)
//	case http.MethodGet:
//		p.showScore(w, r)
//	}
//}

//func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
//	player := strings.TrimPrefix(r.URL.Path, "/players/")
//
//	score := p.store.GetPlayerScore(player)
//
//	if score == 0 {
//		w.WriteHeader(http.StatusNotFound)
//	}
//	fmt.Fprint(w, score)
//}

//func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
//	player := strings.TrimPrefix(r.URL.Path, "/players")
//	p.store.RecordWin(player)
//	w.WriteHeader(http.StatusAccepted)
//}
//
//func (p *PlayerServer) processWin(w http.ResponseWriter) {
//	p.store.RecordWin("Bob")
//	w.WriteHeader(http.StatusAccepted)
//}

//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//	// on verifie la methode de la request
//	if r.Method == http.MethodPost {
//		w.WriteHeader(http.StatusAccepted)
//		return
//	}
//
//	player := strings.TrimPrefix(r.URL.Path, "/players/")
//
//	score := p.store.GetPlayerScore(player)
//
//	if score == 0 {
//		w.WriteHeader(http.StatusNotFound)
//	}
//	fmt.Fprint(w, score)
//}

//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	player := strings.TrimPrefix(r.URL.Path, "/players/")
//
//	score := p.store.GetPlayerScore(player)
//
//	if score == 0 {
//		w.WriteHeader(http.StatusNotFound)
//	}
//	fmt.Fprint(w, score)
//}

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
