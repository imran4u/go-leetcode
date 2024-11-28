package main

import "fmt"

//Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

func main() {
	nums := []int{1, 3, 5, 6}
	// target := 5 // result should be 2
	target := 7 // result should be 4
	// target := 2 // result should be 1

	// nums := []int{1, 3}
	// target := 2 // result should be 1

	fmt.Println(searchInsert(nums, target))

}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	mid := (left + right) / 2

	// run the loop till left is less than right and mid should not go beyound length
	for left <= right && mid < len(nums) {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {

			left = mid + 1

		}
		mid = (left + right) / 2
	}

	return left

}
