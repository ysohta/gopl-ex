package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	clearHTMLFiles()
	retCode := m.Run()
	clearHTMLFiles()
	os.Exit(retCode)
}

func clearHTMLFiles() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".html") {
			os.Remove(f.Name())
		}
	}
}

func TestFetchall(t *testing.T) {
	var tests = []struct {
		urls []string
	}{
		{[]string{"http://gopl.io"}},
		{[]string{"http://gopl.io", "https://golang.org/"}},
	}
	for _, test := range tests {
		fetchall(test.urls)
		cnt := 0
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".html") {
				cnt++
			}
		}
		if cnt != len(test.urls) {
			t.Errorf("number of files differ. Actual:%d Expected:%d", cnt, len(test.urls))
		}

		clearHTMLFiles()
	}
}
