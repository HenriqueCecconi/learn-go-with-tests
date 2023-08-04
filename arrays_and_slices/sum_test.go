package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, wanted %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
	t.Run("sum tails of two slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{0, 9, 2})
		want := []int{7, 11}
		checkSums(t, got, want)
	})
	t.Run("sum tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 2})
		want := []int{0, 11}
		checkSums(t, got, want)
	})
}
