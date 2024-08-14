package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// handleFile()
	// OpenFile()
	// fileReader()
	// readFlow()
	copyFileData()
}

func handleFile() {
	file, err := os.OpenFile("/home/saiddis/Documents/learning_vim.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("file opened")

	_, err = file.Write([]byte("hello from go"))

	if err != nil {
		fmt.Println("writting err:", err)
		return
	}
	fmt.Println("file written successfully")

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

func OpenFile() {
	f, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	fmt.Println("file created")

	f, err = os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	fmt.Println("file opened")

	_, err = f.Write([]byte("hello go"))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data written")

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

func fileReader() {
	// Открытие файла как источник потока данных
	file, err := os.OpenFile("/home/saiddis/Documents/learning_vim.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Чтение файла и вывод его содержимого по частям
	buffer := make([]byte, 32) // Чтение по 1024 байта
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // Достигнут конец файла
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Вывод прочитанной части файла

		fmt.Print(string(buffer[:n]))
	}
}

func readFlow() {
	var input string
	fmt.Print("Enter your name: ")
	fmt.Fscan(os.Stdin, &input)
	fmt.Printf("Hello, %s!\n", input)
}

func copyFileData() {
	srcFile, err := os.OpenFile("/home/saiddis/Documents/learning_vim.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	// Создание файла-назначения для записи
	dstFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dstFile.Close()

	// Копирование данных из файла-источника в файл-назначение
	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Printf("Copied %d bytes from %s to %s.\n", bytesCopied, "source.txt", "destination.txt")
}
