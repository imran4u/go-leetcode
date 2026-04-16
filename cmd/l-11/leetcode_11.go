package main

import "fmt"

// problem : https://leetcode.com/problems/container-with-most-water/

func main() {
	input := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(input))
}

func maxArea(height []int) int {
	area := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := min(height[l], height[r])
		w := r - l
		fmt.Println(h, w, area)
		area = max(area, h*w)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
