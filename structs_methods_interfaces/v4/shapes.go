package main

import "math"

// We can create a simple type using a struct. A struct is just a named collection of fields where you can store data.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of the rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// // Area returns the area of the rectangle.
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
