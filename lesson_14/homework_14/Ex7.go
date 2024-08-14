package main

import (
	"bufio"
	"fmt"
	"os"
)

func readAndWriteToFile(fileName string) error {
	var input string

	fmt.Print("What to write?\n>>> ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	err = writeToFile(fileName, input)
	if err != nil {
		return err
	}

	return nil
}
