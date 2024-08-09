package main

import (
	"errors"
	"reflect"
)

func GetValuesByFunc(i interface{}, f func(v interface{}) bool) ([]interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	var keys []interface{}

	mVal := reflect.ValueOf(i)

	iter := mVal.MapRange()

	for iter.Next() {
		if f(iter.Value().Interface()) {
			keys = append(keys, iter.Key().Interface())
		}
	}

	return keys, nil
}

func ValuesMoreThan1(v interface{}) bool {
	value := reflect.ValueOf(v)
	if value.Type().Kind() != reflect.Int {
		return false
	}

	if value.Int() > 1 {
		return true
	}

	return false
}
