package main

import "strings"

func RemoveDuplicates(str, dup string) string {
	uniqueSymbolIndex := strings.Index(str, dup)
	if uniqueSymbolIndex == -1 {
		return str
	}

	str = ReplaceAll(str, dup, "")
	sPart1 := str[:uniqueSymbolIndex]
	sPart2 := str[uniqueSymbolIndex:]
	return sPart1 + dup + sPart2
}
