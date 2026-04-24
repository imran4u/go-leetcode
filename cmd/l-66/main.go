package main

import "fmt"

func main() {
	digits := []int{2, 4, 9} // result [2,5,0]
	fmt.Println(plusOne(digits))

	fmt.Println("änother way", plusOneA(digits))
}

//Note: see the constrain "1 <= digits.length <= 100"
// it is very large number so you can't hold in int.
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

func plusOneA(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	// all are 9 case
	return append([]int{1}, digits...)
}
