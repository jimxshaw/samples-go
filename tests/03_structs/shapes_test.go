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

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{7.0, 7.0}
		checkArea(t, rect, 49.0)
	})

	t.Run("circles", func(t *testing.T) {
		cir := Circle{10.0}
		checkArea(t, cir, math.Pi*math.Pow(10.0, 2))
	})

}
