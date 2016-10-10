package main

import "fmt"

func main() {
	s := join("foo", " ", "bar")
	fmt.Println(s) // => "foo bar"
}

func join(strs ...string) string {
	concat := ""
	for _, s := range strs {
		concat += s
	}
	return concat
}
