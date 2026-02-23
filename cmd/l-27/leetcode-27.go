package main

import (
	"fmt"
)

func main() {
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// val := 2

	// nums := []int{3,2,2,3}
	// val := 3

	// nums := []int{3, 3}
	// val := 3

	nums := []int{2}
	val := 3

	result := removeElement(nums, val)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	result := nums[:count]
	fmt.Println(result)
	return count

}
