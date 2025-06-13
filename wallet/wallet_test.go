package wallet

import (
	"testing"
)

//var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with founds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient fund", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds) //dispo en global
		assertBalance(t, wallet, Bitcoin(20))
		// si on as une erreur on affiche un message d'erreur et on stop le test
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("didn't get an error but wanted one")
	}

	// permet de convertir l'erreur en string
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
