package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		checkValues(t, got, want)
	})

	t.Run("Circle Perimeter", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83185307179586

		checkValues(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		checkValues(t, got, want)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		checkValues(t, got, want)
	})
}

func checkValues(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got: %g want: %g", got, want)
	}
}
