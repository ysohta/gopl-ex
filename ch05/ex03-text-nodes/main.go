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
		fmt.Fprintf(os.Stderr, "parse failure: %v\n", err)
		os.Exit(1)
	}
	for _, t := range extractText(nil, doc) {
		fmt.Println(t)
	}
}

func extractText(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	texts = extractText(texts, n.FirstChild)
	texts = extractText(texts, n.NextSibling)

	return texts
}
