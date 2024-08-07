package main

import "strings"

func Prefix(str, pref string) bool {
	return strings.HasPrefix(str, pref)
}

func Suffix(str, suff string) bool {
	return strings.HasSuffix(str, suff)
}
