package main

// Rectangle has four sides.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter finds the perimeter of the shape.
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area finds the area of the shape
func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
