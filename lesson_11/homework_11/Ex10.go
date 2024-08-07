package main

import "strings"

func Compare(str1, str2 string) bool {
	if strings.ToLower(str1) == strings.ToLower(str2) {
		return true
	}

	return false
}
