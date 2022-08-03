package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Fount")
  errWordExists = errors.New("That word already exists")
  errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant Delete non-existing word")
)


func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists  
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}