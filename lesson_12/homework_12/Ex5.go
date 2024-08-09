package main

import "strings"

func CountWords(str string) map[string]int {
	wordCount := map[string]int{}
	words := strings.Split(str, " ")

	for _, w := range words {
		wordCount[w] += 1
	}

	return wordCount
}
