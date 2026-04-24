package main

import "fmt"

/*
Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"
Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

Constraints:
1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
	fmt.Println("ïmproved ->", addBinaryImproved(a, b))
}

// main trick to get 1, 0 from byte value to int int(b[n]-'0')
func addBinaryImproved(a string, b string) string {
	m := len(a) - 1
	n := len(b) - 1

	carray := 0
	result := ""
	for m >= 0 || n >= 0 || carray > 0 {
		sum := carray
		if m >= 0 {
			sum += int(a[m] - '0')
			m--
		}

		if n >= 0 {
			sum += int(b[n] - '0')
			n--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carray = sum / 2

	}

	return result
}

// main trick to get 1, 0 from byte value to int int(b[n]-'0')
func addBinary(a string, b string) string {
	m := len(a) - 1
	n := len(b) - 1

	carray := 0
	result := ""
	for m >= 0 && n >= 0 {
		sum := int(a[m]-'0') + int(b[n]-'0') + carray
		result = fmt.Sprintf("%d%s", sum%2, result)
		carray = sum / 2
		m--
		n--
	}
	for m >= 0 {
		sum := int(a[m]-'0') + carray
		result = fmt.Sprintf("%d%s", sum%2, result)
		carray = sum / 2
		m--
	}
	for n >= 0 {
		sum := int(b[n]-'0') + carray
		result = fmt.Sprintf("%d%s", sum%2, result)
		carray = sum / 2
		n--
	}

	if carray == 1 {
		result = fmt.Sprintf("%d%s", 1, result)
	}

	return result
}
