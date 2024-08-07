package main

import "strings"

func ChangeCase(str string, toUpper bool) string {
	if toUpper {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
