package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var in = os.Stdin

func main() {
	doc, err := html.Parse(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, v := range mapElements(map[string]int{}, doc) {
		fmt.Println(k, v)
	}
}

func mapElements(m map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return m
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	m = mapElements(m, n.FirstChild)
	m = mapElements(m, n.NextSibling)

	return m
}
