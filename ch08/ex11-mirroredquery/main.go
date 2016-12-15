package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify URLs")
		os.Exit(1)
	}
	fmt.Println(multiRequest(os.Args[1:]))
}

func multiRequest(urls []string) string {
	var wg sync.WaitGroup
	responses := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			res, err := request(url)
			defer wg.Done()
			if err != nil {
				// error occurs icluding cancellation
				return
			}
			responses <- res
		}(url)
	}
	res := <-responses
	close(done) // cancel the others

	wg.Wait()
	return res // return the quickest response
}

func request(url string) (reponse string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("new request failure: %s", err)
	}

	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetch:%s", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading:%s", err)
	}
	return string(b), nil
}
