package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{5.0, 5.0}
	got := Perimeter(rect)
	want := 20.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{7.0, 7.0}
		got := rect.Area()
		want := 49.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		cir := Circle{10.0}
		got := cir.Area()
		want := math.Pi * 10.0 * 10.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
