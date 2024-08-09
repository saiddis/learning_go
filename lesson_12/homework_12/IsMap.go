package main

import (
	"reflect"
)

func IsMap(i interface{}) bool {
	if reflect.ValueOf(i).Kind() == reflect.Map {
		return true
	}

	return false
}
