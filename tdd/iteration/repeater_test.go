package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("print a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertTests(t, repeated, expected)
	})

	t.Run("print b 6 times", func(t *testing.T) {
		repeated := Repeat("b", 6)
		expected := "bbbbbb"

		assertTests(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 9)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("n", 10))
	// Output: nnnnnnnnnn
}

func assertTests(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%q' but go '%q'", expected, got)
	}
}
