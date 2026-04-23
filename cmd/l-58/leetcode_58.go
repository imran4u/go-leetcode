package main

import "fmt"

func main() {
	input := "Hello World"
	result := lengthOfLastWord(input)
	fmt.Println(result)

	// optimise solution
	fmt.Println(lengthOfLastWordOptimise(input))

}

// This is correct but having one problem, it is scanning the whole string.
// though you need to get the count of last word. here complexcity is O(N)
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0
	temp := 0
	for _, r := range s {
		if r == ' ' {
			if temp != 0 {
				result = temp
			}
			temp = 0
		} else {
			temp++
		}

	}
	if temp != 0 {
		return temp
	}
	return result

}

// Loop from right side and calculate the lenght of last word only
// complexcity is O(K)
func lengthOfLastWordOptimise(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := len(s) - 1
	//step-1: remove all tairling space
	for i >= 0 && s[i] == ' ' {
		i--
	}
	//step-2 : count the word length
	result := 0
	for i >= 0 && s[i] != ' ' {
		result++
		i--
	}
	return result

}
