package structs

import "math"

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
