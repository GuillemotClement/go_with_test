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

// la fonction retourne une erreur
func (w *Wallet) Withdraw(amount Bitcoin) error {
	// on viens test la valeur du wallet avant de faire l'operation de retrait
	if amount > w.balance {
		// throw une erreur si la valeur du compte est inferieur a l'operation
		return errors.New("oh oh")
	}

	w.balance -= amount
	// si ok on rertourne nil qui indique une absence d'erreur
	return nil
}
