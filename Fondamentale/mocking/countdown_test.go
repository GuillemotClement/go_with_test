package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		// on creer un nouveau buffer => une zone memoire temporaire
		buffer := &bytes.Buffer{}
		// on passe ce buffer dnas la fonction => le pointeur et permet a la fonction de modifier le buffer
		Countdown(buffer, &SpyCountdwonOperations{})

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
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdwonOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted call %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
