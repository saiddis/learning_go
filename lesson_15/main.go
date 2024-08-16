package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseYear int    `json:"year"`
	Pages       int    `json:"pages"`
}

func main() {
	book := Book{
		Title:       "Mastery",
		Author:      "Robert Greene",
		ReleaseYear: 2013,
		Pages:       365,
	}

	file, err := os.Create("book.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(&book)

	_, err = file.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	f, err := os.Open("book.json")
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}

	jsonData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("errrr", err)
		return
	}
	var book2 Book

	err = json.Unmarshal(jsonData, &book2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", book2)
}
