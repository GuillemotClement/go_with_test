package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// creer une instance vide de `bytes.Buffer`: bytes.Buffer{}
	// stocke l'adresse memoire de ce buffer avec le &
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
