package arrays

import (
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("array of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum all numbers in some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("sum tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func BenchmarkTestSumMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllMake([]int{1, 2, 3}, []int{4, 5, 6})
	}
}

func BenchmarkTestSumAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllAppend([]int{1, 2, 3}, []int{4, 5, 6})
	}
}
