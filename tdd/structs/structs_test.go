package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shapes, want int) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{10, 40}
		want := 100
		checkPerimeter(t, &r, want)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		want := 62
		checkPerimeter(t, &c, want)
	})
}

func TestArea(t *testing.T) {
	checkAreA := func(t testing.TB, shape Shapes, want int) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{10, 10}
		want := 100
		checkAreA(t, &r, want)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		want := 314
		checkAreA(t, &c, want)
	})
}
