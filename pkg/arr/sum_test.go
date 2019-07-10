package arr

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectSum := func(t *testing.T, expect, sum int, arr []int) {
		t.Helper()
		if expect != sum {
			t.Errorf("got %d want %d given, %v", sum, expect, arr)
		}
	}

	t.Run("5 elems sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		want := 15

		assertCorrectSum(t, want, sum, numbers)
	})

	t.Run("single element array", func(t *testing.T) {
		numbers := []int{1}

		sum := Sum(numbers)
		want := 1

		assertCorrectSum(t, want, sum, numbers)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	assertCorrectSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("two slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertCorrectSum(t, got, want)
	})

	t.Run("empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 6})
		want := []int{0, 15}

		assertCorrectSum(t, got, want)
	})

}
