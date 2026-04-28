package main

import "fmt"

func main() {
	input := []int{3, 2, 3, 1, 1, 1, 4, 4, 4, 4}
	fmt.Println(majorityElement(input))
}

// O(n) and memory O(n)
// func majorityElement(nums []int) int {
//     m := make(map[int]int)
//     for _,v := range nums {
//         m[v]++
//     }
//     max :=0
//     majority := 0
//     for k,v := range m {
//         if v > max {
//             max =v
//             majority = k
//         }
//     }
//     return majority
// }

// complexcity O(n), using Moor's voting algorith, it can give you maximum freq element
// it will calculate the maximum frequency ( voting )element
func majorityElement(nums []int) int {
	num := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			num = v
		}
		if v == num {
			count++
		} else {
			count--
		}

	}

	return num
}
