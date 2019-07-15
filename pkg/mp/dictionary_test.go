package mp

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("kwown word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new", func(t *testing.T) {
		word := "test"
		definition := "test value"

		dict := Dictionary{}
		err := dict.Add(word, definition)

		assertNotError(t, err)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "test value"

		dict := Dictionary{}
		err := dict.Add(word, definition)
		assertNotError(t, err)

		err = dict.Add(word, "other value")
		assertError(t, err, ErrKeyAlreadyExists)
		assertDefinition(t, dict, word, definition)
	})
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	assertNotError(t, err)
	assertStrings(t, got, definition)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("wanted an error '%s' but didnt get one", want)
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNotError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("wanted no error, but got '%s'", got)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s', given '%s'", got, want, "test")
	}
}
