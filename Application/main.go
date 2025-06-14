package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5555", server))
}

//http://localhost:5555/players/Pepper

//func main() {
//	// on definit la fonction appeler lorqu'une requete arrive
//	handler := http.HandlerFunc(PlayerServer)
//	// on lance le serveur et on ecoute sur le port 5000
//	log.Fatal(http.ListenAndServe(":5555", handler))
//}

//	func main() {
//		server := &PlayerServer{}
//		log.Fatal(http.ListenAndServe(":5555", server))
//	}
