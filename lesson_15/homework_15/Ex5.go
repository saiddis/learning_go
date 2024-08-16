package main

import (
	"encoding/json"
	"fmt"
)

func marshalOmittingEmpty(product Product) string {
	b, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
