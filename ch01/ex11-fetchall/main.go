// ex11-fetchall creates files from specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fetchall(os.Args[1:])
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(urls []string) {
	ch := make(chan string)
	for i, url := range urls {
		filepath := fmt.Sprintf("fetched_%d.html", i)
		go fetch(url, ch, filepath)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string, filepath string) {
	// Create a file
	out, err := os.Create(filepath)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer out.Close()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(out, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
