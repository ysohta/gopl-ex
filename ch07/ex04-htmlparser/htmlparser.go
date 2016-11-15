package htmlparser

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func Parse(s string) (*html.Node, error) {
	return html.Parse(*NewReader(s))
}

func NewReader(s string) *io.Reader {
	var r io.Reader
	r = bytes.NewBufferString(s)
	return &r
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
