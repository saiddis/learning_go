package main

import (
	"io"
	"os"
)

func countWords(fileName string) (int, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}

	defer f.Close()
	buf := make([]byte, 128)
	var wordsSum int
	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		for _, v := range buf {
			if string(v) == "\n" || string(v) == " " {
				wordsSum++
			}
		}
	}

	return wordsSum, nil
}
