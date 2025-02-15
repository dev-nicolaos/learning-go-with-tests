package main

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound      = DictionaryErr("No definition found")
	ErrAlreadyExists = DictionaryErr("Tried to add word to dictionary that already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if d[key] == "" {
		return "", ErrNotFound
	}
	return d[key], nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
