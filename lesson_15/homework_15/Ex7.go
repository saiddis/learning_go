package main

import (
	"encoding/json"
	"fmt"
)

func marshalMapInterface(m map[string]interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
