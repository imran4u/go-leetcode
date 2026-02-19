package main

import "fmt"

func main() {
	s := "IV"

	result := convertRomanToDecimal(s)
	fmt.Println(result)
}

func convertRomanToDecimal(s string) int {

	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		current := values[s[i]]

		// Check if next value is greater
		if i+1 < n && current < values[s[i+1]] {
			total -= current
		} else {
			total += current
		}
	}

	return total
}
