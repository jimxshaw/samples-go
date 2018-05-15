package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{2, 4, 6, 8, 10}

	got := Sum(numbers)
	want := 30

	if want != got {
		t.Errorf("got %d want %d give, %v", got, want, numbers)
	}
}
