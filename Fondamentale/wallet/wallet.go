package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	// permet d'utiliser le typage definis dans la struct qui type la valeur
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// permet de modifier le comportement de %s
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// on prepare la valeur du message d'erreur
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds // on passe direct la valeur de l'erreur => sera traiter comme une erreur
	}

	w.balance -= amount
	return nil
}
