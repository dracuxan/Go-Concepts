package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("when supplied a, should return aaaaa", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertTests(t, repeated, expected)
	})

	t.Run("when supplied b, should return bbbbb", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"

		assertTests(t, repeated, expected)
	})
}

func assertTests(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%q' but go '%q'", expected, got)
	}
}
