package main

import (
	"fmt"
)

func main() {
	/* a := 5
	b := 6
	c := a + b
	mt.Println(c) */

	f := sum1(8)
	i := f()
	fmt.Println(i)

	fmt.Println(sumAndMultiply(4, 8))
}

func sum1(a int) func() int {
	b := 6
	return func() int {
		return a + b
	}
}

/* func sum(a, b int) int {
	c := a + b
	return c
}

func printHello(name string) {
	fmt.Println("Hello", name)
}

func inc(a int) {
	a++
	fmt.Println(a)
} */

func sumAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}
