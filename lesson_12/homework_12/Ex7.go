package main

import (
	"errors"
	"reflect"
)

func IsEmpty(i interface{}) (bool, error) {
	if !IsMap(i) {
		return false, errors.New("unsupported type")
	}

	m := reflect.ValueOf(i)
	if len(m.MapKeys()) > 0 {
		return false, nil
	}

	return true, nil
}
