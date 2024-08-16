package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalToMap(jsonData string) map[string]int {
	m := make(map[string]int)
	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return m
}
