package main

import "fmt"

func main() {
	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */

	/* i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	} */

	/* for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	} */

	/* for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(i * j)
		}
		fmt.Println("-------------------------")
	} */

	/* sum := float32(1)
	n := float32(15) */

	// operator := true

	// 10

	/* for i := 1; i <= n; i++ {

		sum += float32(1) / float32(i)
	} */

	//11

	/* for i := 1; i <= n; i++ {
		sum += n ^ 2 + (n + i) ^ 2
	} */

	//12

	/* for i := float32(1); i <= n; i++ {
		sum *= 1 + i/10
	} */

	//13

	/* for i := float32(1); i <= n; i++ {
		sum -= 1 + i/10
		i += 1
		sum += 1 + i/10

	} */

	// series2

	/* var p float64 = 1
	var n float64

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		p = p * n
	}

	fmt.Println(p) */

	var p float64
	var n float64

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)

		p += n - float64(int(n))

	}

	fmt.Println(p)

}
