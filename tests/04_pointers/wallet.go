package main

import "fmt"

type Bitcoin int

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit will deposit money
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance will show the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
