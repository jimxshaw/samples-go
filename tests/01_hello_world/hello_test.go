package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to folks", func(t *testing.T) {
		got := Hello("James")
		want := "Hello, James"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	})

}
