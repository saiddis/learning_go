package main

import (
	"encoding/json"
	"fmt"
)

func marshalMap(m map[string]int) string {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
