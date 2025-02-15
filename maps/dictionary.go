package main

import "errors"
import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", errors.New(fmt.Sprintf("No definition found for %q", key))
	}
	return d[key], nil
}
