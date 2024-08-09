package main

import (
	"errors"
	"reflect"
)

func Update(i interface{}, k interface{}, newV interface{}) error {
	if !IsMap(i) {
		return errors.New("unsupported type")
	}
	m := reflect.ValueOf(i)
	kVal := reflect.ValueOf(k)
	value := reflect.ValueOf(newV)

	m.SetMapIndex(kVal, value)
	return nil
}
