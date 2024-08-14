package main

import (
	"io"
	"os"
)

func readFirstLine(fileName string) (string, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}

	defer f.Close()
	buf := make([]byte, 16)
	var firstLine string

	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		for _, v := range buf {
			if string(v) == "\n" {
				return firstLine, nil
			}
			firstLine += string(v)
		}
	}

	return firstLine, nil
}
