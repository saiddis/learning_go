package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() float64 {
	return float64(r.Width) * float64(r.Height)
}

func main() {
	rect := Rectangle{
		Width:  23,
		Height: 56,
	}
	fmt.Println(rect.Area())
}
