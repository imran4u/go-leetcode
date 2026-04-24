package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println("improve ->", mySqrtImproved(8))
}

// O(log n)
func mySqrtImproved(x int) int {
	if x < 2 {
		return x
	}
	left := 1
	right := x
	mid := left + (right-left)/2
	result := 0

	for left <= right {
		if mid*mid > x {
			right = mid - 1
			mid = left + (right-left)/2
		} else if mid*mid < x {
			result = mid
			left = mid + 1
			mid = left + (right-left)/2
			
		} else {
			return mid
		}
	}

	return result
}

// O(root n)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	for n := 1; n <= x; n++ {
		if n*n > x {
			return n - 1
		}
	}
	return 0
}
