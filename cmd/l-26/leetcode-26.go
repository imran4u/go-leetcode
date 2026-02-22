package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	result := removeDuplicates(nums)
	fmt.Println(result)
}
func removeDuplicates(nums []int) int {
	p := 0
	fp := 0 //forward pointer
	count := len(nums)
	for i := 0; i < count; i++ {
		if nums[fp] > nums[p] {
			p++
			nums[p] = nums[fp]
		}
		fp++
	}
	result := nums[:p+1]
	fmt.Println(result)
	return len(result)
}
