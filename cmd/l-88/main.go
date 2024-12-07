package main

import "fmt"

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3

	merge(num1, m, num2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n <= 0 {
		fmt.Println(nums1)
		return
	}

	i := m + n - 1
	j := m - 1
	k := n - 1

	for ; i >= 0 && k >= 0; i-- {
		fmt.Println("i j k ", i, j, k)

		if j < 0 {
			nums1[i] = nums2[k]
			k--
			continue
		}

		if nums1[j] >= nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}

	}

	fmt.Println(nums1)
}
