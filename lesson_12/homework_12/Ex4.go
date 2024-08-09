package main

import (
	"errors"
	"reflect"
)

func Delete(i interface{}, k interface{}) error {
	if !IsMap(i) {
		return errors.New("unsupported type")
	}

	m := reflect.ValueOf(i)
	kVal := reflect.ValueOf(k)

	m.SetMapIndex(kVal, reflect.Value{})
	return nil
}
