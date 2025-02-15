package main

import "errors"
import "fmt"

func Search(dictionary map[string]string, key string) (string, error) {
	if dictionary[key] == "" {
		return "", errors.New(fmt.Sprintf("No definition found for %q", key))
	}
	return dictionary[key], nil
}
