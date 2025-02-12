package main

import "testing"

func assertMatch(t testing.TB, want, got int, input []int) {
	if got != want {
		t.Errorf("Expected '%d', but got '%d'. Was given '%v'", want, got, input)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 3, 5, 7}

	got := Sum(numbers)
	want := 16

	assertMatch(t, want, got, numbers)
}
