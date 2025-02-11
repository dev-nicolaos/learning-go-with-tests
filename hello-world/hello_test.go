package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Say hello to name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello to world if name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola Juan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jaque", "French")
		want := "Bonjour Jaque"

		assertCorrectMessage(t, got, want)
	})
}
