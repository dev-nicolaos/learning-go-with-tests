package main

import "testing"

func assertError(t testing.TB, expectedErr, realErr error) {
	t.Helper()
	if realErr == nil {
		t.Fatal("Expected an error but recieved nil")
	}

	if expectedErr != realErr {
		t.Errorf("Expected error %s, but recieved %s", expectedErr, realErr)
	}
}

func TestSearch(t *testing.T) {
	t.Run("returns def when available", func(t *testing.T) {

		want := "this is just a test"
		dictionary := Dictionary{"test": want}

		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("Expected no error but got one")
		}

		if got != want {
			t.Errorf("Expected %q, but got %q", want, got)
		}
	})
	t.Run("returns an error when no def", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("foo")

		assertError(t, ErrNotFound, err)
	})
}
