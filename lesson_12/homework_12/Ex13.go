package main

import (
	"errors"
	"reflect"
)

func ConcatMaps(i1, i2 interface{}) (interface{}, error) {
	if !IsMap(i1) || !IsMap(i2) {
		return nil, errors.New("unsupported type")
	}

	mVal1 := reflect.ValueOf(i1)
	mVal2 := reflect.ValueOf(i2)

	if mVal1.Type().Key() != mVal2.Type().Key() ||
		mVal1.Type().Elem() != mVal1.Type().Elem() {
		return nil, errors.New("diffrent map types")
	}

	mType := reflect.MapOf(mVal1.Type().Key(), mVal1.Type().Elem())
	m := reflect.MakeMap(mType)

	iter1 := mVal1.MapRange()
	for iter1.Next() {
		m.SetMapIndex(iter1.Key(), iter1.Value())
	}

	iter2 := mVal2.MapRange()
	for iter2.Next() {
		m.SetMapIndex(iter2.Key(), iter2.Value())
	}

	return m.Interface(), nil
}
