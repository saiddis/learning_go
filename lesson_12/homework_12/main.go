package main

import "fmt"

func main() {
	fmt.Println("---Ex1---")
	m := map[string]int{
		"two":      2,
		"thousand": 0,
		"twenty":   2,
		"four":     4,
	}
	fmt.Println(m)

	fmt.Println("---Ex2---")
	check, err := HasKey(m, "two")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(check)

	fmt.Println("---Ex3---")
	err = Update(m, "four", 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)

	fmt.Println("---Ex4---")
	err = Delete(m, "four")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)

	fmt.Println("---Ex5---")
	str := "hello world"
	fmt.Println(CountWords(str))

	fmt.Println("---Ex6---")
	newM := map[int]string{
		4: "four",
		2: "two",
		1: "one",
	}
	sortedKeys, err := GetSortedKeys(newM)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sortedKeys)

	fmt.Println("---Ex7---")
	check, err = IsEmpty(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(check)

	fmt.Println("---Ex8---")
	inversedM, err := InverseMap(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inversedM)

	fmt.Println("---Ex9---")
	check, err = CheckForDup(m, "two")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(check)

	fmt.Println("---Ex10---")
	values, err := GetValues(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(values...)

	fmt.Println("---Ex11---")
	str = "hello world hello go"
	fmt.Println(CountUniqueWords(str))

	fmt.Println("---Ex12---")
	evenValues, err := GetValuesByCond(m, "even")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(evenValues...)
	fmt.Println("---Ex13---")
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	concatedM, err := ConcatMaps(m, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(concatedM)

	fmt.Println("---Ex14---")
	k, err := GetKeyByVal(m2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(k)

	fmt.Println("---Ex15---")
	sharedPairs, err := GetSharedPairs(m, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sharedPairs)

	fmt.Println("---Ex16---")
	keys, err := GetAllKeys(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys...)

	fmt.Println("---Ex17---")
	strValues, err := GetStringOfValues(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strValues)

	fmt.Println("---Ex18---")
	mCopy, err := GetCopy(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mCopy)

	fmt.Println("---Ex19---")
	slice, err := GetSliceOfPairs(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(slice)

	fmt.Println("---Ex20---")
	keys, err = GetValuesByFunc(m, ValuesMoreThan1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys...)
}
