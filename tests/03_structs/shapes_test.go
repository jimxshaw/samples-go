package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{8.0, 2.0}
		checkPerimeter(t, rect, 20.0)
	})

	t.Run("Circles", func(t *testing.T) {
		cir := Circle{3.0}
		checkPerimeter(t, cir, 2*math.Pi*3.0)
	})

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
