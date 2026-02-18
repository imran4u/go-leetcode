package main

import (
	"fmt"
)

func main() {
	result := isPalindrome(24542)
	fmt.Println("result", result)

}

// This solution is using extra space and loop
// func isPalindrome(x int) bool {
	
// 	arr := []int{}

// 	for x != 0 {
// 		arr = append(arr, x%10)
// 		x = x / 10

// 	}
// 	n := len(arr)
// 	for i := 0; i < n/2; i++ {
// 		if arr[i] == arr[n-i-1] {
// 			continue
// 		}
// 		return false
// 	}

// 	return true
// }

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0

	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// For even length: x == reversedHalf
	// For odd length:  x == reversedHalf/10
	return x == reversedHalf || x == reversedHalf/10
}
