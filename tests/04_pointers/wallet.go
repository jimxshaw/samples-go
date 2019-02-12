package main

import "fmt"

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit will deposit money
func (w Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance will show the balance
func (w Wallet) Balance() int {
	return w.balance
}
