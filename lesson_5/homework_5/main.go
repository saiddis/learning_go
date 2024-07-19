package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("\n---Exersice 1--- \n \n")
	ex_1()
	fmt.Print("\n---Exersice 2--- \n \n")
	ex_2()
	fmt.Print("\n---Exersice 3--- \n \n")
	ex_3()
	fmt.Print("\n---Exersice 4--- \n \n")
	ex_4()
	fmt.Print("\n---Exersice 5--- \n \n")
	ex_5()
	fmt.Print("\n---Exersice 6--- \n \n")
	ex_6()
	fmt.Print("\n---Exersice 7--- \n \n")
	ex_7()
	fmt.Print("\n---Exersice 8--- \n \n")
	ex_8()
	fmt.Print("\n---Exersice 9--- \n \n")
	ex_9()
	fmt.Print("\n---Exersice 10--- \n \n")
	ex_10()
	fmt.Print("\n---Exersice 11--- \n \n")
	ex_11()
	fmt.Print("\n---Exersice 12--- \n \n")
	ex_12()
	fmt.Print("\n---Exersice 13--- \n \n")
	ex_13()
	fmt.Print("\n---Exersice 14--- \n \n")
	ex_14()

}

func ex_1() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func ex_2() {
	for i := 1; i <= 5; i++ {
		i_sqr := i * i
		fmt.Printf("%d squered is %d \n", i, i_sqr)
	}
}

func ex_3() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("3 x %d = %d \n", i, 3*i)
	}
}

func ex_4() {
	var a, b, curr int

	for i := 0; i < 10; i++ {
		if i <= 1 {
			a, curr = i, i
			fmt.Println(curr)
			continue
		}

		curr = a + b
		b = a
		a = curr
		fmt.Println(curr)

	}
}

func ex_5() {
	var a, b int
	fmt.Print("First number >>> ")
	fmt.Scan(&a)
	fmt.Print("Second number >>> ")
	fmt.Scan(&b)

	var result int

	if a%b == 0 || b%a == 0 {

		if a%b == 0 {
			result = b
		} else {
			result = a
		}

		fmt.Println("GCD is", result)
		return
	}

	result = getGCD(a, b, min(a, b))
	fmt.Println("GCD is", result)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getGCD(a, b, div int) int {
	if a%div == 0 && b%div == 0 {
		return div
	}
	return getGCD(a, b, div-1)
}

func ex_6() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		}
	}
}

func ex_7() {
	var num int
	var result bool = true
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	if num < 2 {
		result = false
	}

	for i := 2; i <= num; i++ {
		if num%i == 0 {
			result = false
			if i*2 == num {
				break
			}
		}
	}

	if result {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}

func ex_8() {
	var num int
	fmt.Print("Type a number >>> ")

	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
			if i*2 == num {
				fmt.Println(num)
				break
			}
		}
	}
}

func ex_9() {
	var num float32
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	if num/float32(10) < 1 {
		fmt.Println("Sum of the integers is", num)
		return
	}

	strNum := strconv.Itoa(int(num))
	var sum int

	for _, char := range strNum {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	fmt.Println("Sum of the integers is", sum)
}

func ex_10() {
	var num int
	fmt.Print("To proceed, type a positive number >>> ")
	fmt.Scan(&num)
	for num < 0 {
		fmt.Print("To proceed, type a positive number >>> ")
		fmt.Scan(&num)
	}
}

func ex_11() {
	var num, result int
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	for i := 1; i < num; i++ {
		result += i * (i + 1)
		if result > 1000 {
			break
		}
	}
}

func ex_12() {
	var num int
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	var reversed string
	strNum := strconv.Itoa(int(num))

	for index := len(strNum) - 1; index > -1; index-- {
		reversed += string(strNum[index])
	}

	if reversed == strNum {
		fmt.Println("It is a palindromic number")
	} else {
		fmt.Println("It is not a palindromic number")
	}
}

func ex_13() {
	var num int
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	fmt.Println(buildPyromid(float32(num)))
}

func buildPyromid(height float32) string {
	var pyramid string = ""
	var starsPerFloor float32 = height
	for i := int(height); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			pyramid += " "
		}

		starsPerFloor += 1

		for k := int(starsPerFloor) - i; k > 0; k-- {
			pyramid += "*"
		}

		pyramid += "\n"
	}

	return pyramid
}

func ex_14() {
	var num int
	fmt.Print("Type a number >>> ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		fmt.Printf("---%d---\n \n", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d \n", i, j, i*j)
		}
		fmt.Println()
	}
}
