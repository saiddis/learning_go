package main

import "strings"

type StringInserter interface {
	InsertAt(index int, sub string) string
	InsertAfterWord(word, sub string) string
}

type TextInserter struct {
	s string
}

func (ti TextInserter) InsertAt(index int, sub string) string {
	return ti.s[:index] + sub + ti.s[index:]
}

func (ti TextInserter) InsertAfterWord(word, sub string) string {
	wordLastIndex := strings.LastIndex(ti.s, word)

	return ti.s[:wordLastIndex+1] + " " + sub + " " + ti.s[wordLastIndex:]
}

type Inserter struct {
	StringInserter
}

func NewInserter(si StringInserter) Inserter {
	return Inserter{
		si,
	}
}
