package mapsAndDictionaries

import (
	"errors"
	"testing"
)

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("wouldn't search word:", err)
	}
	if result != definition {
		t.Errorf("got %s, want %s", result, definition)
	}
}
func verifyError(t *testing.T, result error, expected error) {
	t.Helper()
	if !errors.Is(result, expected) {
		t.Errorf("got %s, want %s", result, expected)
	}
}
func stringCompare(t *testing.T, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is an only test"}
	t.Run("search for existing entry", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is an only test"
		stringCompare(t, result, expected)
	})

	t.Run("search for non existing entry", func(t *testing.T) {
		_, err := dictionary.Search("unknown entry")
		verifyError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this an only test"
		err := dictionary.Add(word, definition)
		verifyError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is an only test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		verifyError(t, err, ErrAlreadyExists)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this an only test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		verifyError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this an only test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		verifyError(t, err, ErrNotExistent)

	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this an only test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected word %s, was deleted", word)
	}
}
