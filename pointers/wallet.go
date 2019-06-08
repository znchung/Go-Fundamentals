package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(num Bitcoin) error {
	if num > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= num
	return nil
}

func (w *Wallet) Deposit(num Bitcoin) {
	w.balance += num
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
