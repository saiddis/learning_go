package main

import (
	"errors"
	"reflect"
)

func GetSliceOfPairs(i interface{}) ([][2]interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	var sliceOfPairs [][2]interface{}
	var pair [2]interface{}
	mVal := reflect.ValueOf(i)
	iter := mVal.MapRange()

	for iter.Next() {
		pair[0], pair[1] = iter.Key().Interface(), iter.Value().Interface()
		sliceOfPairs = append(sliceOfPairs, pair)
	}

	return sliceOfPairs, nil
}
