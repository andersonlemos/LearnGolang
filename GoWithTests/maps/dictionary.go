package mapsAndDictionaries

import (
	"errors"
)

type Dictionary map[string]string

type ErrDictionary string

var (
	ErrNotFound      = ErrDictionary("word not found")
	ErrAlreadyExists = ErrDictionary("word already exists")
	ErrNotExistent   = ErrDictionary("word not exists")
)

func (e ErrDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrNotExistent
	case err == nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
