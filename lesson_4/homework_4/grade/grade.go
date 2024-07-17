package grade

import (
	"fmt"
	"os"
	"strconv"
)

func Grage() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a score number:
		0-100`)
		return
	}

	score := os.Args[1]
	intScore, err := strconv.Atoi(score)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(getGrade(intScore))
}

func getGrade(score int) string {
	if score > 0 && score < 31 {
		return "C"
	} else if score > 30 && score < 71 {
		return "B"
	} else if score > 70 && score < 101 {
		return "A"
	} else {
		return `Please provide a score number:
		0-100`
	}
}
