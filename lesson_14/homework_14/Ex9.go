package main

import (
	"io"
	"os"
)

func concatFiles(file1, file2, outputFile string) error {
	f1, err := os.OpenFile(file1, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.OpenFile(file2, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f2.Close()

	outF, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer outF.Close()

	_, err = io.Copy(outF, f2)
	if err != nil {
		return err
	}

	_, err = io.Copy(outF, f1)
	if err != nil {
		return err
	}
	return nil
}
