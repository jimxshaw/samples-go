package main

import (
	"math"
)

// Shape has to have a way to calculate area.
type Shape interface {
	Area() float64
}

// Rectangle must have a width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area finds the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle must have a radius.
type Circle struct {
	Radius float64
}

// Area finds the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter finds the perimeter of the shape.
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}
