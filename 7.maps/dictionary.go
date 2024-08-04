package main

const (
	ErrKeywordNotFound = DictionaryError("could not find the word you were looking for")
	ErrKeywordExists   = DictionaryError("cannot add word because it already exists")
	ErrKeywordNotExist = DictionaryError("cannot update word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	// go map property will return 2 values
	// the second value is a boolean which indicates if the key was found successfully
	definition, ok := d[keyword]

	if !ok {
		return "", ErrKeywordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(keyword, description string) error {
	_, err := d.Search(keyword)

	switch err {
	case ErrKeywordNotFound:
		d[keyword] = description
	case nil:
		return ErrKeywordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(keyword, newDescription string) error {
	_, err := d.Search(keyword)

	switch err {
	case ErrKeywordNotFound:
		return ErrKeywordNotExist
	case nil:
		d[keyword] = newDescription
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(keyword string) {
	// In go, to delete a map, just use `delete` method
	// it takes 2 arguments, the first is the map,
	// and the second is the key to be removed
	delete(d, keyword)
}

func main() {

}
