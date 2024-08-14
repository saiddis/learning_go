package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func replaceWordInFile(fileName, old, new string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	text := string(content)

	modifiedText := strings.ReplaceAll(text, old, new)
	// buf := make([]byte, 128)
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(modifiedText))
	if err != nil {
		return err
	}
	return nil
}
