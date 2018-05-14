package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// Specifying this method is a helper.
		// When the test fails, the line number will be
		// in our function call rather than inside
		// this test helper.
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to folks", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Angie", "Spanish")
		want := "Hola, Angie"

		assertCorrectMessage(t, got, want)
	})

}
