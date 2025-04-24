package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 10}, 40.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTest {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
