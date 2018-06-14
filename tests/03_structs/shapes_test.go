package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{5.0, 5.0}
	got := Perimeter(rect)
	want := 20.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{7.0, 7.0}
	got := Area(rect)
	want := 49.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
