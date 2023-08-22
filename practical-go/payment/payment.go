package main

import (
	"fmt"
	"sync"
)

func main() {
	p := Payment{
		From:   "Joseph",
		To:     "Healthcare Org",
		Amount: 360.16,
	}

	// Payment will only be process once.
	p.Process()
	p.Process()
}

func (p *Payment) Process() {
	// Idempotent operation.
	p.once.Do(p.process)
}

func (p *Payment) process() {
	fmt.Printf("%s -> $%.2f -> %s\n", p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once // Not exported!
}
