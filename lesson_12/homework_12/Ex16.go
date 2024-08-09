package main

import (
	"errors"
	"reflect"
)

func GetAllKeys(i interface{}) ([]interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	var keys []interface{}
	mVal := reflect.ValueOf(i)

	for _, k := range mVal.MapKeys() {
		keys = append(keys, k.Interface())
	}

	return keys, nil
}
