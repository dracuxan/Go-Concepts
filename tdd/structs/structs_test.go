package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	testPeri := []struct {
		name  string
		shape Shapes
		want  int
	}{
		{name: "Rectangle", shape: &Rectangle{a: 12, b: 6}, want: 36},
		{name: "Circle", shape: &Circle{r: 10}, want: 62},
		{name: "Triangle", shape: &Triangle{a: 12, b: 6, c: 2}, want: 20},
	}
	for _, s := range testPeri {
		t.Run(s.name, func(t *testing.T) {
			got := s.shape.Perimeter()
			if got != s.want {
				t.Errorf("got %d want %d", got, s.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testPeri := []struct {
		name  string
		shape Shapes
		want  int
	}{
		{name: "Rectangle", shape: &Rectangle{a: 12, b: 6}, want: 72},
		{name: "Circle", shape: &Circle{r: 10}, want: 314},
		{name: "Triangle", shape: &Triangle{a: 12, b: 6}, want: 36},
	}
	for _, s := range testPeri {
		t.Run(s.name, func(t *testing.T) {
			got := s.shape.Area()
			if got != s.want {
				t.Errorf("got %d want %d", got, s.want)
			}
		})
	}
}
