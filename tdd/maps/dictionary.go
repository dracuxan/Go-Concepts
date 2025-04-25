package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("the word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation because the word doesnot exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(find string) (string, error) {
	if _, ok := d[find]; !ok {
		return "", ErrNotFound
	}

	return d[find], nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = newValue
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
