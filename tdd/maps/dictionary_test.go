package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "just a test"}

	t.Run("basic key search", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "just a test"
		checkKeys(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")
		if got == nil {
			t.Fatal("expected an error")
		}
		checkErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a new key value pair", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "this is just a test"

		dict.Add(key, value)
		checkDefinition(t, dict, key, value)
	})

	t.Run("adding an existing key value", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}

		err := dict.Add(key, "new test")
		checkErrors(t, err, ErrWordExists)
		checkDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key's value", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}
		newValue := "new value"

		err := dict.Update(key, newValue)
		checkErrors(t, err, nil)
		checkDefinition(t, dict, key, newValue)
	})

	t.Run("what if we update a non existing key?", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(key, value)
		checkErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete when a word exist", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "Test defenation"}

		err := dict.Delete(word)
		checkErrors(t, err, nil)

		_, err = dict.Search(word)
		checkErrors(t, err, ErrNotFound)
	})
	t.Run("delete when a word does not exist", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)
		checkErrors(t, err, ErrWordDoesNotExist)
	})
}

func checkDefinition(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	checkKeys(t, got, value)
}

func checkKeys(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func checkErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
