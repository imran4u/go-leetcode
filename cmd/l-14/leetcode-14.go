package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// first input :
	// str := []string{
	// 	"flower", "flow", "flight"
	// }

	str := []string{
		"reflower", "flow", "flight",
	}

	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})

	fmt.Println("sorted, ", str)
	sort.Strings(str)
	fmt.Println(str)
	first := str[0]
	last := str[len(str)-1]
	n := len(first)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		if first[i] != last[i] {
			break
		}
		sb.WriteByte(first[i])
	}
	fmt.Println("Heighest comman factor :", sb.String())
}
