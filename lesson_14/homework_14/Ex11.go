package main

import (
	"bufio"
	"os"
	"strings"
)

func countUniqueWords(fileName string) (int, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	uniqueWords := make(map[string]bool, 0)
	scanner := bufio.NewScanner(f)
	var wordsPerLine []string

	for scanner.Scan() {
		wordsPerLine = strings.Split(scanner.Text(), " ")
		for _, v := range wordsPerLine {
			if ok := uniqueWords[v]; !ok {
				uniqueWords[v] = true
			}
		}
	}

	return len(uniqueWords), nil
}
