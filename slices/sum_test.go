package main

import (
"slices"
"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 3, 5, 7}

	got := Sum(numbers)
	want := 16

	if got != want {
			t.Errorf("Expected %d, but got %d. Was given %v", want, got, numbers)
		}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1,3,5}
	numbers2 := []int{2}
	numbers3 := []int{-5, 10}

	got := SumAll(numbers1, numbers2, numbers3)
	want := []int{9, 2, 5}

	if !slices.Equal(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
