package grade

import (
	"fmt"
)

func Grade() {
	var score int
	fmt.Print(`Type the score number: 0-100 
>>> `)
	fmt.Scan(&score)
	fmt.Println("Your grade:", getGrade(score))
}

func getGrade(score int) string {
	if score > 0 && score < 31 {
		return "C"
	} else if score > 30 && score < 71 {
		return "B"
	} else if score > 70 && score < 101 {
		return "A"
	} else {
		return `Error: given score or input type is not valid`
	}
}
