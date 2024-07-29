package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [8]int{-2, 4, 1, 4, -5, 7, 14, -1}

	fmt.Println("---getMaxNum([8]int) int---")
	fmt.Println(getMaxNum(nums))

	fmt.Println("---getMinNum([8]int) int---")
	fmt.Println(getMinNum(nums))

	fmt.Println("---getPositiveNumsLen([8]int) int---")
	fmt.Println(getPositiveNumsLen(nums))

	fmt.Println("---getSumOfArrayInts([8]int) int---")
	fmt.Println(getSumOfArrayInts(nums))

	fmt.Println("---getAverageNumsOfArray([8]int) int---")
	fmt.Println(getAverageNumOfArray(nums))

	fmt.Println("---deleteNumFromArray([8]int) []int---")
	slice := deleteNumFromArray(nums, 4)
	fmt.Println(slice)

	fmt.Println("---mulArrayNums([8]int, int) [8]int---")
	nums = mulArrayNums(nums, 2)
	fmt.Println(nums)

	fmt.Println("---getNumIndexes([8]int, int) []int---")
	fmt.Println(getNumIndexes(nums, 8))

	fmt.Println("---copyArray([8]int) [8]int---")
	nums2 := copyArray(nums)
	fmt.Println(nums2)

	fmt.Println("---concatArrays([8]int, [8]int) []int---")
	slice = concatArrays(nums, nums2)
	fmt.Println(slice)

	fmt.Println("---swapMinMaxNums([8]int)---")
	swapMinMaxNums(&nums)
	fmt.Println(nums)

	fmt.Println("---isArrPalindrome([8]int) bool---")
	palindromeArr := [8]int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(isArrPalindrome(palindromeArr))

	fmt.Println("---getSecondMaxNum([8]int) int---")
	fmt.Println(getSecondMaxNum(nums))

	fmt.Println("---reverseArr([8]int) [8]int---")
	nums = reverseArr(nums)
	fmt.Println(nums)

	fmt.Println("---deleteDuplicates([8]int) []int---")
	uniqueNums := deleteDuplicates(nums)
	fmt.Println(uniqueNums)
}

func deleteDuplicates(nums [8]int) (slice []int) {
	uniqueNumsMap := map[int]bool{}
	for _, v := range nums {
		if !uniqueNumsMap[v] {
			uniqueNumsMap[v] = true
		}
	}

	for k := range uniqueNumsMap {
		slice = append(slice, k)

	}

	return slice
}

func reverseArr(nums [8]int) [8]int {
	var reversedArr [8]int

	for i, j := 0, len(nums); i < len(nums); i++ {
		reversedArr[i] = nums[j-1]
		j--
	}
	return reversedArr
}

func getSecondMaxNum(nums [8]int) int {
	slice := nums[:]
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	return nums[1]
}

func isArrPalindrome(nums [8]int) bool {
	var reversedArr [8]int

	for i, j := 0, len(nums); i < len(nums); i++ {
		reversedArr[i] = nums[j-1]
		j--
	}

	if reversedArr == nums {
		return true
	}

	return false
}

func swapMinMaxNums(nums *[8]int) {
	maxN := getMaxNum(*nums)
	minN := getMinNum(*nums)
	var maxNIndex, minNIndex int = -1, -1
	for i, v := range nums {
		if maxNIndex != -1 && minNIndex != -1 {
			break
		}

		switch v {
		case maxN:
			maxNIndex = i
		case minN:
			minNIndex = i
		}
	}

	nums[maxNIndex], nums[minNIndex] = nums[minNIndex], nums[maxNIndex]
}

func concatArrays(arr1, arr2 [8]int) []int {
	slice1 := arr1[:]
	slice2 := arr2[:]

	concatenatedSlices := append(slice1, slice2...)
	return concatenatedSlices
}

func copyArray(nums [8]int) [8]int {
	return nums
}

func getNumIndexes(nums [8]int, n int) []int {
	var slice []int
	for i, v := range nums {
		if v == n {
			slice = append(slice, i)
		}
	}

	return slice
}

func mulArrayNums(nums [8]int, n int) [8]int {
	for i, v := range nums {
		nums[i] = v * n
	}

	return nums
}

func deleteNumFromArray(nums [8]int, num int) []int {
	slice := nums[:]

	for i, v := range slice {
		if num == v {
			slice = deleteItem(slice, i)
		}
	}

	return slice
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func getAverageNumOfArray(nums [8]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum / len(nums)
}

func getSumOfArrayInts(nums [8]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

func getPositiveNumsLen(nums [8]int) int {
	var positiveNums []int
	for _, v := range nums {
		if v >= 0 {
			positiveNums = append(positiveNums, v)
		}
	}

	return len(positiveNums)
}

func getMaxNum(nums [8]int) int {
	slice := nums[:]
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	return nums[0]
}

func getMinNum(nums [8]int) int {
	slice := nums[:]
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice[0]
}
