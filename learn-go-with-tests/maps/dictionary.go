package main

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, found := dictionary[word]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dictionary Dictionary) Add(word string, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}
