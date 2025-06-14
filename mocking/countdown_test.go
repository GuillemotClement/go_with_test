package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// on creer un nouveau buffer => une zone memoire temporaire
	buffer := &bytes.Buffer{}
	// on passe ce buffer dnas la fonction => le pointeur et permet a la fonction de modifier le buffer
	Countdown(buffer)
	// en interne, le buffer contiens des []byte => c'est un tableau de bytes
	// avec .String() on convertis le []byte en string
	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
