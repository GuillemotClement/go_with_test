package wallet

type Wallet struct {
	balance int
}

// dans une fonction, on passe une copie de la struct
// la methode ne modifie pas la valeur de la struct
// w est une copie de la struct
func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
