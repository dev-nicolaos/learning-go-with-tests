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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
}

func assertMatch(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	t.Run("returns def when available", func(t *testing.T) {
		want := "this is just a test"
		dictionary := Dictionary{"test": want}

		got, err := dictionary.Search("test")

		assertNoError(t, err)
		assertMatch(t, want, got)
	})
	t.Run("returns an error when no def", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("foo")

		assertError(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "foo"
		definition := "bar"
		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertMatch(t, definition, dictionary[word])
	})

	t.Run("existing word", func(t *testing.T) {
		word := "foo"
		dictionary := Dictionary{word: "bar"}
		newDefinition := "baz"
		err := dictionary.Add(word, newDefinition)

		assertError(t, ErrAlreadyExists, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "foo"
		dictionary := Dictionary{word: "bar"}
		newDefinition := "baz"

		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertMatch(t, newDefinition, dictionary[word])
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("foo", "bar")

		assertError(t, ErrNotFound, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "foo"
		dictionary := Dictionary{word: "bar"}

		deleteErr := dictionary.Delete(word)

		assertNoError(t, deleteErr)

		_, searchErr := dictionary.Search(word)

		assertError(t, ErrNotFound, searchErr)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("foo")

		assertError(t, ErrNotFound, err)
	})

}
