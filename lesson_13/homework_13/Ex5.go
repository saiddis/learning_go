package main

type StringConcatenator interface {
	Concat() string
	Join(sep string) string
}

type StringJoiner struct {
	s [2]string
}

func (sj StringJoiner) Concat() string {
	return sj.s[0] + sj.s[1]
}

func (sj StringJoiner) Join(sep string) string {
	return sj.s[0] + sep + sj.s[1]
}

type ArrayHandler struct {
	StringConcatenator
}

func NewArrayHandler(sc StringConcatenator) ArrayHandler {
	return ArrayHandler{
		sc,
	}
}
