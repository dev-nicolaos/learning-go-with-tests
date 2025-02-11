package main

import "testing"

const basicFailTemplate = "got %q want %q"


func TestHello(t *testing.T) {
	t.Run("Say hello to name", func(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		t.Errorf(basicFailTemplate, got, want)
	}
	})
	t.Run("Say hello to world if name is empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf(basicFailTemplate, got, want)
		}
	})
}
