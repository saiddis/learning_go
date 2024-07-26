package main

import (
	"fmt"
)

func main() {
	// slice1 := []int{1, 2, 3}
	// slice2 := slice1
	// slice2 = append(slice2, 3)
	// slice1[0] = 10
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	slice := [7]int{1, 2, 3, 4, 5, 6, 7}
	iterateOverArrayInt(slice)

	slice2 := [5]string{"a", "b", "c", "d", "e"}
	iterateOverArraystring(slice2)

	slice3 := [4]bool{true, false, true, false}
	iterateOverArrayBool(slice3)

	var slice4 [10]int
	for _, v := range slice4 {
		fmt.Println(v)
	}

	for i := 1; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}

	slice5 := [3]bool{true, false, false}
	slice5[0] = false

	slice6 := [5]string{"i", "am", "writing", "a", "loop"}
	for _, v := range slice6 {
		fmt.Println(v)
	}

	slice7 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range slice7 {
		if slice7[0] < v && v < slice7[9] {
			fmt.Println(v)
		}
	}
}

func iterateOverArrayBool(slice [4]bool) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func iterateOverArraystring(slice [5]string) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func iterateOverArrayInt(slice [7]int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}
