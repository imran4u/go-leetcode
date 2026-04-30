package main

import "fmt"

func main() {
	s := "aabb"
	fmt.Println(firstUniqChar(s))
	fmt.Println(firstUniqChar_Opt(s))
}

// Time O(N) and space O(1) (since max 26 chars)
func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}

	return -1

}

// Time O(N) and space O(1) (since max 26 chars)
func firstUniqChar_Opt(s string) int {
	freq := [26]int{} // Note better to use freq array over the map

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
