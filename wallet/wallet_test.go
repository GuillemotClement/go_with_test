package wallet

import "testing"

func TestWallet(t *testing.T) {
	// wallet prend une struct
	wallet := Wallet{}
	// methode de la struct
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
