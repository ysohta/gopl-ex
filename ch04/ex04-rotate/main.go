package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	rotate(s, -2)
	fmt.Printf("%v\n", s)
	// =>"[2 3 4 5 0 1]"

	rotate(s, 4)
	fmt.Printf("%v\n", s)
	// =>"[4 5 0 1 2 3]"
}

func rotate(s []int, shift int) {

	var n = shift % len(s)
	if n < 0 {
		n = (n + len(s)) % len(s)
	}

	// bubble rotate
	for i := 0; i < n; i++ {
		for j := len(s) - 1; j > 0; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}
