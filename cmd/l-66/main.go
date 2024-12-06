package main

import "fmt"

func main() {
	digits := []int{2, 4, 9} // result [2,5,0]
	fmt.Println(plusOne(digits))
}

//Note: see the constrain "1 <= digits.length <= 100"
// it is very large number so you can hold in int.
// so iterate from right and check for 9, if there is no 9, just increment and return
// in case of all 9, ex 99, append 1 from begning and then return.

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] == 9 {
			digits[i] = 0
			continue
		}

		digits[i]++
		return digits

	}

	digits = append([]int{1}, digits...)

	return digits
}
