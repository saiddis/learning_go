package main

import "strings"

func DeleteAll(str, subStr string) string {
	return strings.ReplaceAll(str, subStr, "")
}
