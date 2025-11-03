package structs

import "math"

type Rectangle struct {
	a int
	b int
}

type Circle struct {
	r int
}

type Shapes interface {
	Perimeter() int
	Area() int
}

func (r *Rectangle) Area() int {
	return r.a * r.b
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

func (c *Circle) Area() int {
	ans := math.Pi * float64(c.r*c.r)
	return int(ans)
}

func (c *Circle) Perimeter() int {
	ans := 2 * math.Pi * float64(c.r)
	return int(ans)
}
