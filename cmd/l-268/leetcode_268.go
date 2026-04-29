package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}

	fmt.Println("O(N^2) -> ", missingNumber(nums))

	fmt.Println("O(N) & O(N) -> ", missingNumberOpt1(nums))
	fmt.Println("O(N) -> ", missingNumberOpt2(nums))

	fmt.Println("O(N) using Xor-> ", missingNumberOpt3(nums))
}


// O(N^2)
func missingNumber(nums []int) int {

	for i := 0; i <= len(nums); i++ {
		found := false
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				found = true
				break
			}

		}
		if !found {
			return i
		}
	}
	return 0
}

// O(N) & O(N)
func missingNumberOpt1(nums []int) int {
	n := len(nums)
	arr := make([]int, n+1)
	for _, v := range nums {
		arr[v] = -1
	}
	for i, v := range arr {
		if v == 0 {
			return i
		}
	}
	return 0
}

// O(N)
func missingNumberOpt2(nums []int) int {
	// sum := 0
	
	// why len(nums)+1 , because max value will be len(nums)+1
	// for i := range len(nums)+1 {
	// 	sum += i
	// }

	n := len(nums)
	 sum := n * (n + 1) / 2

	for _, v := range nums {
		sum -= v
	}
	return sum
}

//best solution
func missingNumberOpt3(nums []int) int {
	
    n := len(nums)
    xor := n

    for i, v := range nums {
        xor ^= i ^ v
    }

    return xor
}

