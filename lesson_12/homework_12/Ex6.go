package main

import (
	"errors"
	"reflect"
	"sort"
)

func GetSortedKeys(i interface{}) (interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	mVal := reflect.ValueOf(i)
	var keys []interface{}

	iter := mVal.MapRange()
	for iter.Next() {
		keys = append(keys, iter.Key().Interface())
	}
	sort.Slice(keys, func(i, j int) bool {
		if i > j {
			return true
		}
		return false
	})

	return keys, nil
}
