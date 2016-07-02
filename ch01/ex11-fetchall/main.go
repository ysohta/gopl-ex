// ex11-fetchall creates loaded html files.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var out io.Writer = os.Stdout

func main() {
	start := time.Now()
	fetchall(os.Args[1:])
	fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Fprintln(out, <-ch)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2d %7d %s", secs, nbytes, url)
}
