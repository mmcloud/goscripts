package dictionary

type Dictionary map[string]string

const (
	ErrNotFound    = DictError("could not find the key you were looking for")
	ErrKeyFound    = DictError("key already found in dictionary")
	ErrKeyNotFound = DictError("could not update key as already defined")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

func (dict Dictionary) Search(word string) (string, error) {
	value, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)
	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return ErrKeyFound
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyNotFound
	case nil:
		dict[key] = value
	default:
		return err
	}
	return nil
}

func (dict Dictionary) Delete(key string) {
	delete(dict, key)
}
