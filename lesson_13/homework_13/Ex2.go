package main

import "strings"

type Formatter interface {
	ToUpper() string
	ToLower() string
}

type MyFormatter struct {
	s string
}

func (mf MyFormatter) ToUpper() string {
	return strings.ToUpper(mf.s)
}

func (mf MyFormatter) ToLower() string {
	return strings.ToLower(mf.s)
}

type Format struct {
	Formatter
}

func NewFormatter(fmtr Formatter) Format {
	return Format{
		fmtr,
	}
}
