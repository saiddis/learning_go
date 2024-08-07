package main

import "strings"

func ReplaceFirst(str, oldStr, newStr string) string {
	return strings.Replace(str, oldStr, newStr, 1)
}
