package main

import "fmt"

func main() {
	//NIL map painic : assignment to entry in nil map
	// var m map[string]string
	// m["a"] = "b"
	// fmt.Println(m)

	var m map[string]string
	m = make(map[string]string)
	m["a"] = "b"
	fmt.Println(m)

	m2 := make(map[string]string)
	m2["a"] = "c"
	fmt.Println(m2)

	var s []string
	s = append(s, "a")
	fmt.Println(s)

}
