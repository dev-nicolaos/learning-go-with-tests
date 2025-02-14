package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficentFunds = errors.New("Insufficent funds")

type Litcoin int

func (l Litcoin) String() string {
	return fmt.Sprintf("%d LTC", l)
}

type Wallet struct {
	balance Litcoin
}

func (w *Wallet) Deposit(amount Litcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Litcoin) error {
	if amount > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Litcoin {
	return w.balance
}
