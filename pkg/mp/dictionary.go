package mp

import "errors"

var ErrNoKeyFound = errors.New("could not find the key")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNoKeyFound
	}
	return value, nil
}
