package main

import "fmt"

type Bitcoin int

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

// Withdraw will take out money
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance will show the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
