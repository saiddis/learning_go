package main

import (
	"io"
	"os"
)

func copyFile(src, dst string) error {
	from, err := os.OpenFile(src, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}
