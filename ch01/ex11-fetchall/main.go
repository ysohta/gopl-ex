// ex11-fetchall prints time to fetch URLs in CSV.
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

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	start := time.Now()
	for _, csv := range os.Args[1:] {
		fetchFromCsv(csv)
	}
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

	var cnt, nData int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		records := strings.Split(scanner.Text(), ",")
		// take second column as URL
		go fetch(records[1], ch, ioutil.Discard)
		nData++
	}
	for {
		select {
		case <-done:
			fmt.Println("canceled")
			return
		case msg := <-ch:
			fmt.Println(msg)
			cnt++
			if cnt == nData {
				return
			}
		}
	}
}

func fetch(url string, ch chan<- string, out io.Writer) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		close(done) // broadcast cancel
		return
	}

	nbytes, err := io.Copy(out, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		close(done) // broadcast cancel
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
