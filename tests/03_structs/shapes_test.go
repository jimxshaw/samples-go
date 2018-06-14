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

	t.Run("Triangles", func(t *testing.T) {
		tri := Triangle{3.0, 4.0, 5.0, 2.4}
		checkPerimeter(t, tri, 12.0)
	})

}

func TestArea(t *testing.T) {

	// Declare a slice of anonymous structs with two
	// fields. It's filled with cases.
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{3.0, 4.0}, 12.0},
		{"Circle", Circle{4.0}, math.Pi * math.Pow(4.0, 2)},
		{"Triangle", Triangle{3.0, 4.0, 5.0, 2.4}, 6.0},
	}

	// Iterate over each item in the slice, using
	// the struct fields to run our tests.
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}

}
