package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Q: What is the most common word (ignore case) in sherlock.txt?
// Word frequency

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()
	wordFrequency(file)

	mapDemo()

	// `s` is a "raw" string where \ is just a \
	// rather than escape character like \n.
	path := `C:\to\new\folder\report.csv`
	fmt.Println(path)
}

func mapDemo() {
	var stocks map[string]float64 // symbol -> count
	sym := "GOOG"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	// stocks[sym] = 483.27 // cannot do this before initializing the map first
	stocks = make(map[string]float64)
	stocks[sym] = 483.27
	stocks["AAPL"] = 363.85

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k := range stocks { // keys
		fmt.Println(k)
	}

	for k, v := range stocks { // keys & values
		fmt.Println(k, " -> ", v)
	}

	for _, v := range stocks { // values
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	// fmt.Printf("%p\n", &stocks) // must use Printf for actual memory address
	delete(stocks, "AAPL") // no panic
}

// Variables will execute before main.
// "What's your name?" -> [What s your name]
var wordRegex = regexp.MustCompile(`[a-zA-Z]+`)

// The init function will also execute before main.
// func init() {}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count

	for s.Scan() {
		words := wordRegex.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
