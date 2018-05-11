package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Go"

	if got != want {
		t.Errorf("got '%s' but want '%s'", got, want)
	}
}
