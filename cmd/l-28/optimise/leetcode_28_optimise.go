package main

import "fmt"

func main() {
	//haystack := "sadbutsad"
	//needle := "sad"

	// haystack := "leetcode"
	// needle := "leeto"

	haystack := "abc"
	needle := "c"

	index := getFirstIndex(haystack, needle)
	fmt.Println(index)

}

func getFirstIndex(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// internally == will take O(m), so solution is still O(m*n)
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}
