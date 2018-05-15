package main

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()

		if want != got {
			t.Errorf("got %d want %d give, %v", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)
		want := 12

		assertCorrectMessage(t, got, want, numbers)
	})

}
