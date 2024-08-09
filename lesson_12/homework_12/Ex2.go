package main

import (
	"errors"
	"reflect"
)

func HasKey(i interface{}, k interface{}) (bool, error) {
	if !IsMap(i) {
		return false, errors.New("unsupported type")
	}

	m := reflect.ValueOf(i)
	kVal := reflect.ValueOf(k)

	if m.MapIndex(kVal).IsValid() {
		return true, nil
	}

	return false, nil
}
