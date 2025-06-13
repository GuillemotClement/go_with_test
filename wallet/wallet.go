package wallet

import "fmt"

type Wallet struct {
	balance int
}

// dans une fonction, on passe une copie de la struct
// la methode ne modifie pas la valeur de la struct
// w est une copie de la struct
// pour modifier la valeur dans la struct il faut passer la struct avec un pointer
func (w *Wallet) Deposit(amount int) {
	// %p permet d'afficher la notation hexa
	fmt.Printf("adress of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
