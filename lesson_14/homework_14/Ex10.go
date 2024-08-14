package main

import "os"

func fileExists(fileName string) bool {
	f, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}
