package main

import "strings"

type DuplicateRemover interface {
	RemoveDuplicates() string
	RemoveDuplicatesCaseInsensitive() string
}

type TextDuplicateRemover struct {
	s string
}

func (tdr TextDuplicateRemover) RemoveDuplicates() string {
	uniqueWords := make(map[string]bool, 0)
	var wordsSequence []string
	words := strings.Split(tdr.s, " ")

	for _, v := range words {
		if b := uniqueWords[v]; !b {
			uniqueWords[v] = true
			wordsSequence = append(wordsSequence, v)
		}
	}

	return strings.Join(wordsSequence, " ")
}

func (tdr TextDuplicateRemover) RemoveDuplicatesCaseInsensitive() string {
	uniqueWords := make(map[string]bool, 0)

	var wordsSequence []string
	words := strings.Split(strings.ToLower(tdr.s), " ")

	for _, v := range words {
		if b := uniqueWords[v]; !b {
			uniqueWords[v] = true
			wordsSequence = append(wordsSequence, v)
		}
	}

	return strings.Join(wordsSequence, " ")
}

type DuplicateHandler struct {
	DuplicateRemover
}

func NewDuplicateHandler(dr DuplicateRemover) DuplicateHandler {
	return DuplicateHandler{
		dr,
	}
}
