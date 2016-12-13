package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	depth int
)

type link struct {
	url   string
	depth int
}

func crawl(l link) []link {
	fmt.Println(l.depth, l.url)
	list, err := Extract(l.url)
	if err != nil {
		log.Print(err)
	}

	links := make([]link, len(list))
	dpth := l.depth + 1
	for _, c := range list {
		links = append(links, link{c, dpth})
	}
	return links
}

func init() {
	flag.IntVar(&depth, "depth", 3, "depth")
	flag.Parse()
}

func main() {
	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		links := make([]link, len(flag.Args()))
		for _, arg := range flag.Args() {
			links = append(links, link{arg, 1})
		}
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
			if !seen[l.url] && l.depth <= depth {
				seen[l.url] = true
				unseenLinks <- l
			}
		}
	}
}
