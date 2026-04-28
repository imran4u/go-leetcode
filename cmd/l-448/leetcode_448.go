package main

import "fmt"

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
	Input: nums = [4,3,2,7,8,2,3,1]
	Output: [5,6]
Example 2:
	Input: nums = [1,1]
	Output: [2]


*/

func main() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// input :=[]int{1,1}
	fmt.Println("O(n) and memory O(1) : ", findDisappearedNumbers(input))
	fmt.Println("optimise to O(n) and memory O(1) : ", findDisappearedNumbersOptimal(input))

}

func findDisappearedNumbersOptimal(nums []int) []int {
	result := make([]int, 0)

	for _, v := range nums {
		// convert value to any unique index by -1
		if v < 0 { // it is alreay visited value.
			v = -v
		}
		if nums[v-1] > 0 { //+ve value means not check, so mark it as -ve(you can mark other things as well)
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 { // value is positive that means it has not marked above
			result = append(result, i+1) // add one in index as number start from 1, and index from 0
		}
	}
	return result
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	m := make(map[int]struct{})
	result := make([]int, 0)
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			result = append(result, i)
		}

	}
	return result
}
