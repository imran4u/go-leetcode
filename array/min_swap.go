package array

/**

Note: Hyqoo dated 15 Feb test
find minimum number of swap to sort an array.

**/

import (
	"fmt"
	"sort"
)

type pair struct {
	value int32
	index int
}

func minSwap(arr []int32) int32 {
	n := len(arr)
	pairs := make([]pair, n)

	// Step 1: Store value and original index
	for i := 0; i < n; i++ {
		pairs[i] = pair{value: arr[i], index: i}
	}

	// Step 2: Sort by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	visited := make([]bool, n)
	swaps := int32(0)

	// Step 3: Detect cycles
	for i := 0; i < n; i++ {

		// Already visited or already in correct position
		if visited[i] || pairs[i].index == i {
			continue
		}

		cycleSize := int32(0)
		j := i

		for !visited[j] {
			visited[j] = true
			j = pairs[j].index
			cycleSize++
		}

		if cycleSize > 1 {
			swaps += cycleSize - 1
		}
	}

	return swaps

}

func main() {

	input := []int32{
		3, 4, 1, 2,
	}
	fmt.Println(input)
	count := minSwap(input)
	fmt.Println("count", count)

}
