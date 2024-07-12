package main

import (
	"fmt"
	"math"
)

const p float32 = 3.14
const a int = 3
const b int = 6
const c int = 5

func exercise_1() {
	P := 4 * a
	fmt.Println("1)", P)
}

func exercise_2() {
	S := a
	fmt.Println("2)", S)
}

func exercise_3() {
	P := 2 * (a + b)
	S := a * b
	fmt.Println("3)", P, S)
}

func exercise_4() {
	d := 6
	L := p * float32(d)
	fmt.Println("4)", L)
}

func exercise_5() {
	V := a ^ 3
	S := a * 6
	fmt.Println("5)", V, S)
}

func exercise_6() {
	V := a * b * c
	S := 2 * (a*b + b*c + a*c)
	fmt.Println("6)", V, S)
}

func exercise_7() {
	R := 8
	L := float32(2) * p * float32(R)
	S := p * float32(R)
	fmt.Println("7)", L, S)
}

func exercise_8() {
	average := (a + b) / 2
	fmt.Println("8)", average)
}

func exercise_9() {
	mul := a * b
	sqrt := math.Sqrt(float64(mul))
	fmt.Println("9)", sqrt)
}

func exercise_10() {
	a_pow := a ^ 2
	b_pow := b ^ 2
	sum := a_pow + b_pow
	diff := a_pow - b_pow
	mul := a_pow * b_pow
	div := a_pow / b_pow
	fmt.Println("10)", sum, diff, mul, div)
}

func exercise_11() {
	L := 88
	m := float32(L) / float32(100)
	fmt.Println("11)", m)
}

func exercise_12() {
	M := 14578
	T := float32(M) / float32(1000)
	fmt.Println("12)", T)
}

func exercise_13() {
	b := 5425
	kb := float32(b) / float32(1024)
	fmt.Println("13)", kb)
}

func exercise_14() {
	A := 90
	B := 15
	Bs_in_A := A / B
	fmt.Println("14)", Bs_in_A)
}

func exercise_15() {
	A := 345
	B := 154
	remained_length := A % B
	fmt.Println("15)", remained_length)
}

func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
	exercise_7()
	exercise_8()
	exercise_9()
	exercise_10()
	exercise_11()
	exercise_12()
	exercise_13()
	exercise_14()
	exercise_15()
}
