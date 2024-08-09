package main

import (
	"errors"
	"reflect"
)

func GetKeyByVal(i interface{}, v interface{}) (interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	val := reflect.ValueOf(v)
	mVal := reflect.ValueOf(i)
	if val.Type() != mVal.Type().Elem() {
		return nil, errors.New("incompatible parameter types were given")
	}
	var key interface{}
	iter := mVal.MapRange()

	for iter.Next() {
		if reflect.DeepEqual(val.Interface(), iter.Value().Interface()) {
			key = iter.Key().Interface()
			break
		}

	}

	return key, nil
}
