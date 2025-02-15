package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("No definition found")
var ErrAlreadyExists = errors.New("Tried to add word to dictionary that already exists")

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", ErrNotFound
	}
	return d[key], nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch (err) {
		case ErrNotFound:
			d[word] = definition
		case nil:
			return ErrAlreadyExists
		default:
			return err
	}
	return nil
}
