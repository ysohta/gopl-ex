package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// usage: ./ex09-expand 'foo $bar baz'
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
	fmt.Println(expand(s, strings.ToUpper))
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "$") {
			words[i] = f(strings.TrimPrefix(w, "$"))
		}
	}
	return strings.Join(words, " ")
}
