package structs

import "math"

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Perimeter() float64 {
	return 2 * circle.Radius * math.Pi
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
