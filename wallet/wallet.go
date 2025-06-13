package wallet

import "fmt"

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

func (w *Wallet) Withdraw(amout Bitcoin) {
	w.balance -= amout
}
