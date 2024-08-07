package main

func Reverse(str string) string {
	var reversed string
	for i := Len(str) - 1; i >= 0; i-- {
		reversed += str[i : i+1]
	}

	return reversed
}
