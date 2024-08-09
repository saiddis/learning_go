package main

import (
	"errors"
	"reflect"
)

func GetSharedPairs(i1, i2 interface{}) (interface{}, error) {
	if !IsMap(i1) || !IsMap(i2) {
		return nil, errors.New("unsupported type")
	}

	mVal1 := reflect.ValueOf(i1)
	mVal2 := reflect.ValueOf(i2)

	if mVal1.Type().Key() != mVal2.Type().Key() ||
		mVal1.Type().Elem() != mVal2.Type().Elem() {
		return nil, errors.New("diffrent map types")
	}

	mType := reflect.MapOf(mVal1.Type().Key(), mVal1.Type().Elem())
	sharedPairs := reflect.MakeMap(mType)

	for _, k := range mVal1.MapKeys() {
		for _, k2 := range mVal2.MapKeys() {
			if reflect.DeepEqual(k.Interface(), k2.Interface()) &&
				reflect.DeepEqual(mVal1.MapIndex(k).Interface(), mVal2.MapIndex(k2).Interface()) {
				sharedPairs.SetMapIndex(k, mVal1.MapIndex(k))
			}

		}
	}

	return sharedPairs.Interface(), nil
}
