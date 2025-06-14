package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// objet qui implemente `io.Writter` dans un buffer fournis
	buffer := bytes.Buffer{}
	// on passe le buffer (la ou la fonction doit ecrire
	Greet(&buffer, "Chris")
	// apres l'appel, on viens lire ce qui a ete ecrit dans le buffer
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
