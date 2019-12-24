package pointer

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance int
}

var ErrNotEnoughBalance = errors.New(fmt.Sprintf("not enough balance to withdraw"))

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int)  {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) error {
	if w.balance < amount {
		return ErrNotEnoughBalance
	}

	w.balance -= amount
	return nil
}