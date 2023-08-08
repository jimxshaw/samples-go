package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 10)
	banner("G☺", 10)

	s := "G☺"
	fmt.Println("len: ", len(s))
	// code point = rune ~= unicode character
	// This for range loop iterate by rune.
	for i, r := range s {
		fmt.Println(i, r)
		// rune (int32)
		fmt.Printf("%c of type %T\n", r, r)
	}

	// byte (uint8)
	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y) // x=1, y=1
	// Use %#v in debug/log because it outputs the type info.
	fmt.Printf("x=%#v, y=%#v\n", x, y) // x=1, y="1"

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("g☺g", isPalindrome("g☺g"))
	fmt.Println("g☺g☺", isPalindrome("g☺g☺"))
	fmt.Println("gogo", isPalindrome("gogo"))
}

// isPalindrome("g") -> true
// isPalindrome("go") -> flase
// isPalindrome("gog") -> true
// isPalindrome("g☺g") -> true
// isPalindrome("g☺g☺") -> false
// isPalindrome("gogo") -> false
func isPalindrome(s string) bool {
	// Convert string to slice of runes to ensure the
	// function works for unicode as well.
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	padding := (width - utf8.RuneCountInString(text)) / 2
	// This for loop iterate byte by byte.
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
