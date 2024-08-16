package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func encodeToFile(i interface{}, fileName string) {
	IKind := reflect.ValueOf(i).Type().Kind()
	if IKind != reflect.Map && IKind != reflect.Struct &&
		IKind != reflect.Slice {
		fmt.Println("unsupported type")

	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(i)
	if err != nil {
		fmt.Println(err)
	}
}
