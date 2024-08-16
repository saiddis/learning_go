package main

import (
	"encoding/json"
	"fmt"
)

func marshalPerson(person Person) string {
	b, err := json.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
