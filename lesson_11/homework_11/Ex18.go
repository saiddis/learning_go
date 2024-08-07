package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DecodeHTMLSymbols(str string) string {
	firstCodeIndex := strings.Index(str, "&")
	lastCodeIndex := strings.LastIndex(str, ";")
	if firstCodeIndex == -1 || lastCodeIndex == -1 {
		return str
	}

	htmlCode := str[firstCodeIndex : lastCodeIndex+1]
	codeNum := str[firstCodeIndex+2 : lastCodeIndex]
	utf8index, err := strconv.Atoi(codeNum)
	if err != nil {
		fmt.Println(err)
		return str
	}

	return strings.Replace(str, htmlCode, string(utf8index), 1)
}
