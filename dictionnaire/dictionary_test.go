package dictionnaire

import "testing"

func TestSearsh(t *testing.T) {
	dictionary := map[string]string{"test": "this is juste a test"}
	got := Search(dictionary, "test")
	want := "this is juste a test"

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
