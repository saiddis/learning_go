package main

import (
	"bufio"
	"os"
)

func reverseReadFile(fileName string) (string, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var fileLines []string
	var reversedFile string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text()+"\n")
	}

	for i := len(fileLines) - 1; i >= 0; i-- {
		reversedFile += fileLines[i]
	}

	return reversedFile, nil
}
