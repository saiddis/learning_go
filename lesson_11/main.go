package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	// str := "важно"
	// b := str[0]
	// fmt.Println(b)

	// str2 := "HEllo world"
	// sub := str2[6:11]
	// fmt.Println(sub)
	// a := "世界"
	// aBytes := []byte(a)
	// fmt.Println(a[0:3])
	// fmt.Println(aBytes)

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(i, a[i])
	// }

	// r, s := utf8.DecodeRune(aBytes)
	// fmt.Println(r, s)

	// r, s = utf8.DecodeRuneInString(a)
	// fmt.Println(r, s)

	// a = "a,b,c,d"
	// split := strings.Split(a, ",")
	// fmt.Println(split)
	//
	strConv()
	useUnicode()
	useUtf8()
	fmt.Println(getNextUTF8Symbol("a"))
}

func strConv() {
	s := "42"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	i = 42
	s = strconv.Itoa(i)
	fmt.Println(s)

	s = "3.14"

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	f = 3.14
	s = strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println(s)
}

func useUnicode() {
	r := 'A'
	fmt.Println(unicode.IsLetter(r))
	r = '1'
	fmt.Println(unicode.IsDigit(r))

	s := "Hello, World!"
	upper := toUpper(s)
	fmt.Println(upper)

	r = 'A'
	fmt.Println(unicode.IsLetter(r))
	fmt.Println(unicode.IsDigit(r))
}

func toUpper(s string) string {
	result := []rune{}
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		result = append(result, unicode.ToUpper(r))
		s = s[size:]
	}
	return string(result)
}

func useUtf8() {
	s := "Hello, 世界"
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("%c ", r)
		s = s[size:]
	}

	s = "Hello, 世界"
	count := utf8.RuneCountInString(s)
	fmt.Printf("Number of runes: %d\n", count)

	s = "Hello, 世界"
	valid := utf8.ValidString(s)
	fmt.Printf("string valid %t\n", valid)

	r := '世'
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Printf("Encoded: % x\n", buf)
	decodedRune, _ := utf8.DecodeRune(buf)
	fmt.Printf("Decoded rune: %c\n", decodedRune)
}

func getNextUTF8Symbol(s string) string {
	buf := make([]byte, 2)
	r, _ := utf8.DecodeRuneInString(s)
	utf8.EncodeRune(buf, r)
	buf[0] += 1
	r, _ = utf8.DecodeRune(buf)

	return string(r)
}
