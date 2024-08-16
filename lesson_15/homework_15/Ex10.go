package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalIgnoringNil(jsonData string) Person3 {
	var person Person3
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println(err)
		return Person3{}
	}

	return person
}
