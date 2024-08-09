package main

import (
	"errors"
	"reflect"
)

func GetCopy(i interface{}) (interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	return reflect.ValueOf(i), nil
}
