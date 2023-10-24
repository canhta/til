package main

import "fmt"

func main() {

	s := []string{"a", "b", "c", "d", "e", "f"}

	l := s[2:5]
	fmt.Println(l) // [c d e]

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println(l) // [a b c d e]

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println(l) // [c d e f]
}
