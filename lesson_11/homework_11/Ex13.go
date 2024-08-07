package main

func CountSymbols(str, sym string) map[string]int {
	symCount := map[string]int{}

	for _, r := range str {
		symCount[string(r)] += 1
	}

	return symCount
}
