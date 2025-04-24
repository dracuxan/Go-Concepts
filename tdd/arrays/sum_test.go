package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("for array [1,2,3,4,5] sum = 15", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15
		assertTest(t, got, want, nums)
	})

	t.Run("for array [1,4,5,8,1] sum = 19", func(t *testing.T) {
		nums := []int{1, 4, 5, 8, 1, 10}

		got := Sum(nums)
		want := 29
		assertTest(t, got, want, nums)
	})
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{5, 6, 5, 3, 1})
	}
}

func assertTest(t testing.TB, got, want int, nums []int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d' given %#v", got, want, nums)
	}
}
