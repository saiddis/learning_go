package main

import (
	"strconv"
	"unicode"
)

func SumStrNums(str string) int {
	var sum int
	for _, v := range str {
		if unicode.IsDigit(v) {
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				continue
			}
			sum += digit
		}
	}

	return sum
}
