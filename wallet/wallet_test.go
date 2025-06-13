package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	// wallet prend une struct
	wallet := Wallet{}
	// methode de la struct
	wallet.Deposit(10)

	got := wallet.Balance()
	// & permet d'acceder a l'adresse memoire de la valeur de wallet.balance
	fmt.Printf("adresse of balance in test is %p \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
