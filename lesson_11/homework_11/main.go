package main

import (
	"fmt"
)

func main() {
	str1 := "Hello, World!"
	str2 := "!"

	fmt.Println("---Ex1---")
	fmt.Println(Concat(str1, str2))

	fmt.Println("---Ex2---")
	fmt.Println(Len(str1))

	fmt.Println("---Ex3---")
	fmt.Println(Contains(str1, str2))

	fmt.Println("---Ex4---")
	fmt.Println(FirstPos(str1, str2))

	fmt.Println("---Ex5---")
	fmt.Println(ReplaceAll(str1, str2, "?"))

	fmt.Println("---Ex6---")
	fmt.Println(TrimSpaces(str1))

	fmt.Println("---Ex7---")
	fmt.Println(ChangeCase(str1, true))

	fmt.Println("---Ex8---")
	fmt.Println(Reapeat(str1, 3))

	fmt.Println("---Ex9---")
	fmt.Println(Split(str1, ","))

	fmt.Println("---Ex10---")
	fmt.Println(Compare(str1, str1))

	fmt.Println("---Ex11---")
	fmt.Println(ReplaceFirst(str1, "World", "Go"))

	fmt.Println("---Ex12---")
	fmt.Println(Reverse(str1))

	fmt.Println("---Ex13---")
	fmt.Println(CountSymbols(str1, str2))

	fmt.Println("---Ex14---")
	fmt.Println(DeleteAll(str1, str2))

	fmt.Println("---Ex15---")
	fmt.Println(Count(str1, "World"))

	fmt.Println("---Ex16---")
	fmt.Println(Prefix(str1, "Hello"))
	fmt.Println(Suffix(str1, "World"))

	fmt.Println("---Ex17---")
	fmt.Println(RemoveDuplicates(Reapeat(str2, 3), str2))

	fmt.Println("---Ex18---")
	htmlCode := "<p>Hello, World&#33;</p>"
	decodedHtml := DecodeHTMLSymbols(htmlCode)
	fmt.Println(decodedHtml)

	fmt.Println("---Ex19---")
	var dashedStr DashString = "Im writing code"
	fmt.Println(dashedStr.Transform())

	fmt.Println("---Ex20---")
	strNums := "12345"
	fmt.Println(SumStrNums(strNums))

	fmt.Println("---Ex21---")
	fmt.Println(InversePos(str1))

	fmt.Println("---Ex22---")
	fmt.Println(CountUniqueWords("Hello, Hello, Hello, World!"))
}
