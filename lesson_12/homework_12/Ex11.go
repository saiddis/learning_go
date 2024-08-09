package main

import "strings"

func CountUniqueWords(str string) int {
	splittedStr := strings.Split(str, " ")
	uniqueWords := map[string]bool{}

	for _, v := range splittedStr {
		uniqueWords[v] = true
	}

	return len(uniqueWords)
}
