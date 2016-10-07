package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	breadthFirst(crawl, os.Args[1:], save)
}

func breadthFirst(f func(item string) []string, worklist []string, post func(rawurl string) error) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
				if err := post(item); err != nil {
					fmt.Errorf("post :%v\n", err)
				}
			}
		}
		fmt.Printf("%v\n", worklist)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extractSameDomain(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func save(rawurl string) error {
	filename, _, err := fetch(rawurl)
	if err != nil {
		return err
	}

	fmt.Printf("filenmame=%s\n", filename)

	return nil
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	// create dir
	dir := "." + strings.TrimSuffix(resp.Request.URL.Path, "/")
	dir = strings.TrimSuffix(dir, local)
	fmt.Printf("dir:%s\nfile:%s\n", dir, local)
	fmt.Println("create dir:", dir)
	if err = os.MkdirAll(dir, 0777); err != nil {
		return
	}

	f, err := os.Create(path.Join(dir, local))
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
