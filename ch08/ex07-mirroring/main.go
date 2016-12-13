package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var (
	host string
	in   string
	out  string
)

func crawl(l string) []string {
	fmt.Println(l)
	list, err := Extract(l)
	if err != nil {
		log.Print(err)
	}

	links := make([]string, len(list))
	for _, c := range list {
		links = append(links, c)
	}
	return links
}

func init() {
	flag.StringVar(&in, "in", "http://www.gopl.io", "input URL")
	flag.Parse()
}

func main() {
	u, err := url.Parse(in)
	if err != nil {
		fmt.Printf("invalid link:%s\n", err)
		os.Exit(1)
	}
	host = u.Host

	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		links := []string{in}
		worklist <- links
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for l := range unseenLinks {
				foundLinks := crawl(l)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, l := range list {
			if !seen[l] && isSameHost(l) {
				seen[l] = true

				unseenLinks <- l
			}
		}
	}
}

func isSameHost(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	return u.Host == host
}

// func mirrorContent(url string) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		return fmt.Errorf("getting %s: %s", url, resp.Status)
// 	}

// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		return fmt.Errorf("parsing %s as HTML: %v", url, err)
// 	}
// }
