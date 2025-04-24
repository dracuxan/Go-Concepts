package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{1, 2, 3, 4}, hasPerimeter: 8.0},
	}

	for _, tt := range perimeterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("got %#v want %g", got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6, 3, 4}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
