package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		if err := fetch(url); err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	}
}

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch: %v", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("fetch: reading %s: %v", url, err)
	}
	fmt.Fprintf(out, "%s", b)
	return nil
}
