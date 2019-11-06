package main

import (
	"errors"
	"fmt"
)

// Bitcoin is an int type
type Bitcoin int

// Stringer is an interface
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit will put in money
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// ErrInsufficientFunds is a global variable
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw will take out money
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance will show the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
