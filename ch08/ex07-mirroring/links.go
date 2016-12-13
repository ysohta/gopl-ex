package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func Extract(rawURL string) ([]string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", rawURL, resp.Status)
	}

	// doc, err := html.Parse(resp.Body)
	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", rawURL, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}

				links = append(links, fmt.Sprintf("%s://%s%s", link.Scheme, link.Host, link.Path))

				// replace to local path
				if link.Host == host {
					a.Val = fmt.Sprintf("%s%s", link.Host, link.Path)
					fmt.Println(a.Val)
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	// save file
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	path := u.Path
	if path == "" {
		path = "/"
	}

	filename := fmt.Sprintf("%s%s", u.Host, path)
	if strings.HasSuffix(filename, "/") {
		filename += "index.html"
	}
	f, err := createFile(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	html.Render(f, doc)

	return links, nil
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

func createFile(filename string) (*os.File, error) {
	path := filepath.Dir(filename)
	if err := os.MkdirAll(path, 0777); err != nil {
		return nil, err
	}

	return os.Create(filename)
}
