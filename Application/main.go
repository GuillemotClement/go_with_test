package main

import (
	"log"
	"net/http"
)

func main() {
	// on definit la fonction appeler lorqu'une requete arrive
	handler := http.HandlerFunc(PlayerServer)
	// on lance le serveur et on ecoute sur le port 5000
	log.Fatal(http.ListenAndServe(":5555", handler))
}
