package main

type SubstringExtractor interface {
	Substr(start, length int) string
	Suffix(length int) string
}

type TextExtractor struct {
	s string
}

func (te TextExtractor) Substr(start, length int) string {
	return te.s[start : length+start]
}

func (te TextExtractor) Suffix(length int) string {
	return te.s[:length]
}

type Extractor struct {
	SubstringExtractor
}

func NewExtractor(se SubstringExtractor) Extractor {
	return Extractor{
		se,
	}
}
