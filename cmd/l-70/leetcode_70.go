package main

import "fmt"

/**
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
	Input: n = 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps
Example 2:
	Input: n = 3
	Output: 3
	Explanation: There are three ways to climb to the top.
	1. 1 step + 1 step + 1 step
	2. 1 step + 2 steps
	3. 2 steps + 1 step

Constraints:
1 <= n <= 45
**/

func main(){
fmt.Println(climbStairs(5))
}

// number of way to climb n stair will be the sum of way to climb n-1 + n-2
func climbStairs(n int) int {

	if n <=2 {
		return n
	}
	ways := make([]int, n+1)
	ways[1] =1
	ways[2] = 2
	for i:=3;i<=n;i++ {
		ways[i]=ways[i-1]+ ways[i-2]
	}
    return ways[n]
}