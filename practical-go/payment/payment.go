package main

import "fmt"

func main() {
	p := Payment{
		From:   "Joseph",
		To:     "Healthcare Org",
		Amount: 360.16,
	}

	p.Process()
}

func (p *Payment) Process() {
	fmt.Printf("%s -> $%.2f -> %s\n", p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD
}
