package main

import (
	"errors"
	"reflect"
)

func CheckForDup(i interface{}, k interface{}) (bool, error) {
	if !IsMap(i) {
		return false, errors.New("unsupported type")
	}
	set := map[interface{}]bool{}

	v := reflect.ValueOf(i)
	mapType := reflect.MapOf(v.Type().Key(), v.Type().Elem())
	m := reflect.MakeMap(mapType)

	for _, k := range m.MapKeys() {
		set[m.MapIndex(k)] = true
	}

	if m.Len() == len(set) {
		return true, nil
	}

	return false, nil
}
