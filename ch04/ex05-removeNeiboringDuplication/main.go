package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "c", "c", "c"}
	s = removeNeiboringDuplication(s)

	fmt.Println(s)
	// =>"[a b c]"
}

func removeNeiboringDuplication(s []string) []string {
	if len(s) == 0 {
		return s
	}

	out := s[:1]
	for i := 1; i < len(s); i++ {
		last := out[len(out)-1]
		if s[i] != last {
			out = append(out, s[i])
		}
	}
	return out
}
