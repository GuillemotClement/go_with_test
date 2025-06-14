package dictionnaire

import "testing"

func TestSearsh(t *testing.T) {
	dictionary := Dictionary{"test": "this is juste a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is juste a test"

		assertStrings(t, got, want)

	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknow")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatalf("expected to get an error")
		}
		// fonction Error() pour afficher une string pour le message d'erreur
		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
