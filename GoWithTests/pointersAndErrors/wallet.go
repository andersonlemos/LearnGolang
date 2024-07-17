package pointersAndErrors

import (
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

var withdrawError = func(balance Bitcoin, amount Bitcoin) error {
	return fmt.Errorf("you cannot withdraw from %s to %s", balance.String(), amount.String())
}

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", int(*b))
}
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return withdrawError(w.balance, amount)
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
