package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
}

// Variables will execute before main.
var wordRegex = regexp.MustCompile(`[a-zA-Z]+`)

// The init function will also execute before main.
// func init() {}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	lnum := 0

	for s.Scan() {
		lnum++
		s.Text() // current line
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	fmt.Println("num lines: ", lnum)

	return nil, nil
}
