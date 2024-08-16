package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func decodeJSONFile(fileName string) Person {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return Person{}
	}
	defer f.Close()

	var person Person

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println(err)
		return Person{}
	}

	return person
}
