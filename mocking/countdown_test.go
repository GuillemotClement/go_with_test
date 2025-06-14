package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// on creer un nouveau buffer => une zone memoire temporaire
	buffer := &bytes.Buffer{}
	// on ajoute l'espion pour connaitre le nombre d'appel
	spySleeper := &SpySleeper{}
	// on passe ce buffer dnas la fonction => le pointeur et permet a la fonction de modifier le buffer
	Countdown(buffer, spySleeper)
	// en interne, le buffer contiens des []byte => c'est un tableau de bytes
	// avec .String() on convertis le []byte en string
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
