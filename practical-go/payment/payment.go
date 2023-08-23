package main

import (
	"fmt"
	"sync"
	"time"
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
	t := time.Now()

	// Idempotent operation.
	// Anonymous function must be used because
	// Do argument must be parameterless.
	p.once.Do(func() { p.process(t) })
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s\n", ts, p.From, p.Amount, p.To)
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once // Not exported!
}
