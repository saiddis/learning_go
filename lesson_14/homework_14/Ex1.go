package main

import (
	"io"
	"os"
)

func CountCharacters(fileName string) (int, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	buf := make([]byte, 128)
	var sum int
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		sum += n
	}

	return sum, nil
}
