package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient fund", func(t *testing.T) {
		// on place 20 bitcoin dans le portefeuille
		startingBalance := Bitcoin(20)
		// on initalise le wallet a 20
		wallet := Wallet{startingBalance}
		// on viens catch un cas d'erreur
		// on essaie de retirer sans du wallet alors qu'il a 20, cela provoque une erreur catcher
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		// si on as une erreur on affiche un message d'erreur et on stop le test
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})

}
