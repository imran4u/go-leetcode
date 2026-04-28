package main

import "fmt"

/*
Valid anagram :

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true
Example 2:
	Input: s = "rat", t = "car"
	Output: false

*/

func main() {
	s := "anagramm"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		m[r]--
		//early exit here it is may not required here.
		if m[r] < 0 { // for anagram it should be always >=0,
			return false
		}
	}

	// This is not requried as len is going to check first.
	// for _, v := range m {
	// 	if v!=0 {
	// 		return  false
	// 	}
	// }

	return true
}
