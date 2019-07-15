package mp

const (
	ErrKeyNotFound      = DictionaryErr("could not find the key")
	ErrKeyAlreadyExists = DictionaryErr("key already exists")
	ErrKeyDoesNotExists = DictionaryErr("key does not exist")
)

type (
	DictionaryErr string
	Dictionary    map[string]string
)

// DictionaryErr implements 'error' interface
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrKeyDoesNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
