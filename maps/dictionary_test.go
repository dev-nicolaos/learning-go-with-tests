package main

import "testing"

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error but recieved nil")
	}
}

func TestSearch(t *testing.T) {
	t.Run("returns def when available", func(t *testing.T) {

		want := "this is just a test"
		dictionary := map[string]string{"test": want}

		got, err := Search(dictionary, "test")

		if err != nil {
			t.Fatal("Expected no error but got one")
		}

		if got != want {
			t.Errorf("Expected %q, but got %q", want, got)
		}
	})
	t.Run("returns an error when no def", func(t *testing.T) {
		dictionary := map[string]string{}
		_, err := Search(dictionary, "foo")

		assertError(t, err)
	})
}
