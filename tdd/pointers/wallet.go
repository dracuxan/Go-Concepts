// module to implement a functionality of bitcoin wallet
package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance += ammount
}

func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if w.balance < ammount {
		return ErrInsufficientFunds
	}

	w.balance -= ammount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
