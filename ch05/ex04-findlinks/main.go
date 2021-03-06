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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
