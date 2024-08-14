package main

import (
	"fmt"
)

func main() {
	fmt.Println("---Ex1---")
	n, _ := CountCharacters("new_file.txt")
	fmt.Println(n)
	fmt.Println("---Ex2---")
	lines, _ := countLines("new_file.txt")
	fmt.Println(lines)

	fmt.Println("---Ex3---")
	words, _ := countWords("new_file.txt")
	fmt.Println(words)

	fmt.Println("---Ex4---")
	writeToFile("new_file.txt", "hello from go\n")

	fmt.Println("---Ex5---")
	firstLine, _ := readFirstLine("new_file.txt")
	fmt.Println(firstLine)

	fmt.Println("---Ex6---")
	copyFile("new_file.txt", "../destination.txt")
	newFileWords, _ := countWords("../destination.txt")
	fmt.Println(newFileWords)

	fmt.Println("---Ex7---")
	err := readAndWriteToFile("../example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	firstLine, _ = readFirstLine("../example.txt")
	fmt.Println(firstLine)

	fmt.Println("---Ex8---")
	reversedFile, _ := reverseReadFile("new_file.txt")
	fmt.Println(reversedFile)

	fmt.Println("---Ex9---")
	err = concatFiles("new_file.txt", "../destination.txt", "concated.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---Ex10---")
	fmt.Println(fileExists("new_file.txt"))

	fmt.Println("---Ex11---")
	n, _ = countUniqueWords("new_file.txt")
	fmt.Println(n)

	fmt.Println("---Ex12---")
	replaceWordInFile("new_file.txt", "hello", "bye")

}
