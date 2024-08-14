package main

import "os"

func writeToFile(fileName, content string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
