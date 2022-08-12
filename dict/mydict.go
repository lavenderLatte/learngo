package mydict

import "errors"

// var errNotFound = errors.New("key not found")
// var errExist = errors.New("key already exist")

// can write above two lines like below as well
var (
	errNotFound  = errors.New("key not found")
	errExist     = errors.New("key already exist")
	errCanDelete = errors.New("key not found; can't delete")
)

// this is not a struct, it just a type Dictionary
// for example, type Money is integer
type Money int

// "Dictionary" is alias for map
type Dictionary map[string]string

// METHOD on TYPE
// go can have methods for not only struct but also for types
func (d Dictionary) Search(key string) (string, error) {
	// d[key] returns two things: value, bool
	value, exist := d[key]
	if exist {
		return value, nil // key is in the map
	}
	return "", errNotFound // key is not in the map
}

func (d Dictionary) Add(key string, val string) error {
	// _, exist := d[key]

	// let's use already implemented function
	_, exist := d.Search(key)
	if exist == errNotFound {
		d[key] = val // adding new key-val pair to dictionary
		return nil
	}
	return errExist

	// using switch-case
	// _, exist := d.Search(key)
	// switch exist {
	// case errNotFound:
	// 	d[key] = val
	// case nil:
	// 	return errExist
	// }
	// return nil
}

func (d Dictionary) Update(key, val string) error {
	_, exist := d.Search(key)
	if exist == errNotFound {
		return errNotFound
	}
	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, exist := d.Search(key)
	if exist == nil {
		delete(d, key)
		return nil
	}
	return errCanDelete
}
