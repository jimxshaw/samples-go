package main

import (
	"math"
)

// Shape has to have a way to calculate area.
type Shape interface {
	Area() float64
	Perimeter() float64
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

// Perimeter finds the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle must have a radius.
type Circle struct {
	Radius float64
}

// Area finds the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter finds the circumference of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
