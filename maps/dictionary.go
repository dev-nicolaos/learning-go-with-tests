package main

import "errors"
import "fmt"

type Dictionary map[string]string

var ErrNotFound = errors.New(fmt.Sprintf("No definition found"))

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", ErrNotFound
	}
	return d[key], nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
