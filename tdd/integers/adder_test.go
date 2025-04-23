package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("sum 2,2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertTest(t, sum, expected)
	})

	t.Run("sum 9,2", func(t *testing.T) {
		sum := Add(9, 2)
		expected := 11

		assertTest(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output:4
}

func assertTest(t testing.TB, got int, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%d' got '%d'", expected, got)
	}
}
