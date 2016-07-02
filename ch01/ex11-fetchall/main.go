// ex11-fetchall creates files from specified URL.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fetchFromCsv("top-1m.csv")
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchFromCsv(csvpath string) {
	f, err := os.Open(csvpath)
	if err != nil {
		fmt.Errorf("cannot open file: %s", err)
		return
	}
	defer f.Close()

	ch := make(chan string)

	scanner := bufio.NewScanner(f)
	var cnt int
	for scanner.Scan() {
		records := strings.Split(scanner.Text(), ",")
		go fetch(records[1], ch)
		cnt++
	}
	for i := 0; i < cnt; i++ {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
