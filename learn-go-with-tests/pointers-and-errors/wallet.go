package wallet

import "errors"
import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
