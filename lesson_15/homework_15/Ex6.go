package main

import (
	"encoding/json"
	"fmt"
)

func marshalEmbeddingStructs(person Person2) string {
	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func unmarshalEmbeddingStructs(jsonData string) Person2 {
	var person Person2

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println(err)
		return Person2{}
	}

	return person
}
