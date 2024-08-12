package main

import (
	"strings"
)

type StringReverser interface {
	Reverse() string
	ReverseWords() string
}

type TextReverser struct {
	s string
}

func (tr TextReverser) Reverse() string {
	var reversed []byte

	for i := len(tr.s) - 1; i >= 0; i-- {
		reversed = append(reversed, tr.s[i])
	}

	return string(reversed)
}

func (tr TextReverser) ReverseWords() string {
	var reversedWords string
	words := strings.Split(tr.s, " ")

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords += words[i]
	}

	return reversedWords
}

type Reverser struct {
	StringReverser
}

func NewReverser(sr StringReverser) Reverser {
	return Reverser{
		sr,
	}
}
