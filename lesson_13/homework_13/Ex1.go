package main

import "strings"

type StringProcessor interface {
	Length() int
	WordCount() int
}

type MyString struct {
	s string
}

func (ms MyString) Length() int {
	return len(ms.s)
}

func (ms MyString) WordCount() int {
	words := strings.Split(ms.s, " ")

	return len(words)
}

type String struct {
	StringProcessor
}

func NewString(sp StringProcessor) String {
	return String{
		sp,
	}
}
