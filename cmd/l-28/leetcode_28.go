package main

import "fmt"

/**
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

**/

//Solution : O(m*n)

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

	length := len(haystack)

	for i := 0; i < length; i++ {
		matched := false
		for j := 0; j < len(needle); j++ {
			// if i+j < length {
			// 	fmt.Println(i, j, haystack[i+j], needle[j])
			// }

			if i+j < length && haystack[i+j] == needle[j] {
				if j == len(needle)-1 {
					matched = true
					break
				}
			} else {
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}
