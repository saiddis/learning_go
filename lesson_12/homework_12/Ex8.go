package main

import (
	"errors"
	"reflect"
)

func InverseMap(i interface{}) (interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	v := reflect.ValueOf(i)
	invertedMapType := reflect.MapOf(v.Type().Elem(), v.Type().Key())
	invertedMap := reflect.MakeMap(invertedMapType)

	for _, k := range v.MapKeys() {
		v := v.MapIndex(k)
		invertedMap.SetMapIndex(v, k)
	}

	return invertedMap.Interface(), nil
}
