package main

import "fmt"

func main() {

	// nums := []int{2, 7, 11, 15}
	// target := 9
	// ans = [0,1]

	nums := []int{3, 2, 4}
	target := 6
	// Output: [1,2]

	result := twoSum(nums, target)
	fmt.Println("result", result)

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[target-nums[i]] = i
	}

	result := make([]int, 2)

	for index, value := range nums {

		mI, ok := m[value]
		if ok {

			if mI < index {
				result[0] = mI
				result[1] = index
				break
			} else if mI > index {
				result[0] = index
				result[1] = mI
				break
			}

		}

	}

	return result
}
