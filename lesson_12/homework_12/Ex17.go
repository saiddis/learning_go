package main

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func GetStringOfValues(i interface{}) (string, error) {
	if !IsMap(i) {
		return "", errors.New("unsupported type")
	}

	var values []string
	var str string
	mVal := reflect.ValueOf(i)

	iter := mVal.MapRange()

	for iter.Next() {
		str = strconv.Itoa(int(iter.Value().Int()))
		values = append(values, str)
	}

	return strings.Join(values, ","), nil
}
