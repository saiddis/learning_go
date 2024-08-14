package main

import (
	"io"
	"os"
)

func countLines(fileName string) (int, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}

	defer f.Close()
	buf := make([]byte, 128)
	var linesSum int
	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		for _, v := range buf {
			if string(v) == "\n" {
				linesSum++
			}
		}
	}

	return linesSum, nil
}
