package main

import "strings"

type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}

type TextCleaner struct {
	s string
}

func (tc TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(tc.s)
}

func (tc TextCleaner) RemoveSpaces() string {
	return strings.TrimFunc(tc.s, func(r rune) bool {
		if string(r) == " " {
			return true
		}
		return false
	})
}

type Trimmer struct {
	StringCleaner
}

func NewTrimmer(sc StringCleaner) Trimmer {
	return Trimmer{
		sc,
	}
}
