package main

import "strings"

func InversePos(str string) string {
	var inversed string
	splittedStr := strings.Split(str, " ")

	for i := len(splittedStr) - 1; i >= 0; i-- {
		inversed += splittedStr[i] + " "
	}

	return inversed
}
