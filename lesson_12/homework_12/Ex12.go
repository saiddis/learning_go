package main

import (
	"errors"
	"reflect"
)

func GetValuesByCond(i interface{}, cond string) ([]interface{}, error) {
	if !IsMap(i) {
		return nil, errors.New("unsupported type")
	}

	mVal := reflect.ValueOf(i)
	iter := mVal.MapRange()
	var keys []interface{}

	switch cond {
	case "odd":
		for iter.Next() {
			if iter.Value().Int()%2 != 0 {
				keys = append(keys, iter.Key())
			}
		}
		return keys, nil

	case "even":
		for iter.Next() {
			if iter.Value().Int()%2 == 0 {
				keys = append(keys, iter.Key())
			}
		}
		return keys, nil

	default:
		return nil, errors.New("supported conditions are 'odd' and 'even'")
	}
}
