package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// affecte la struct a la variable
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
