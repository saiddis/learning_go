package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalPerson(jsonData string) Person {
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println(err)
		return Person{}
	}

	return person
}
