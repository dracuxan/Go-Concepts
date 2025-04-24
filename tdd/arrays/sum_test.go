package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("for array [1,2,3,4,5] sum = 15", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15
		assertTest(t, got, want, nums)
	})

	t.Run("for array [1,8,1,10] sum = 20", func(t *testing.T) {
		nums := []int{1, 8, 1, 10}

		got := Sum(nums)
		want := 20
		assertTest(t, got, want, nums)
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	got := SumAll([]int{1, 2}, []int{0, 2})
	want := []int{3, 2}
	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Testing non empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{3, 9})
		want := []int{9, 9}
		checkSums(t, got, want)
	})
	t.Run("Testing for empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
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
