package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch: %v", err)
	}
	_, err = io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("fetch: reading %s: %v", url, err)
	}
	return nil
}
