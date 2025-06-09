package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")  //on veut tester la fonction
	want := "Hello, Chris" // le resultat attevdu de la foncton tester

	// si le resultat de la fonction est different alors on affiche le message d'erreur
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
