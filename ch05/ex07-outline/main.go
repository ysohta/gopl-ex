package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int
var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		err := printOutline(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "printOutline: %v\n", err)
			continue
		}
	}
}

func printOutline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		s := n.Data
		for _, a := range n.Attr {
			s += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}

		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s<%s>\n", depth*2, "", s)
		} else {
			fmt.Fprintf(out, "%*s<%s/>\n", depth*2, "", s)
		}
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Fprintf(out, "%*s<!--%s-->\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		trimmed := strings.TrimSpace(n.Data)
		if trimmed != "" {
			fmt.Fprintf(out, "%*s%s\n", depth*2, "", trimmed)
		}
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
