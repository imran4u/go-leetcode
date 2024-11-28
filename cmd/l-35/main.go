package main

import "fmt"

//Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5 // result should be 2
	// target := 7 // result should be 4
	// target := 2 // result should be 1

	// nums := []int{1, 3}
	// target := 2 // result should be 1

	fmt.Println(searchInsert(nums, target))
	// Alternate way to solve the same problem
	fmt.Println(searchInsertAlt1(nums, target))

}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	// run the loop till left is less than right
	for left <= right {
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

// Alternate way using recursive
func searchInsertAlt1(nums []int, target int) int {
	v := 0
	len := len(nums)

	if nums[0] > target {
		return 0
	}

	if nums[len-1] < target {
		return len
	}

	mid := len / 2
	v = mid
	if nums[mid] > target {
		v = searchInsertAlt1(nums[:mid], target)
	}

	if nums[mid] < target {
		v += searchInsertAlt1(nums[mid:], target)
	}

	return v

}
