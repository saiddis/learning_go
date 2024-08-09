package main

import (
	"errors"
	"reflect"
)

func GetValues(i interface{}) ([]interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	var values []interface{}
	m := reflect.ValueOf(i)
	iter := m.MapRange()

	for iter.Next() {
		values = append(values, iter.Value().Interface())
	}

	return values, nil
}
