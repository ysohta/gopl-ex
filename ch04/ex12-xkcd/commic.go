package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const path = "./archives"

type Commic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Url        string
}

func search(query string) {
	files := listFiles(path)

	for _, r := range files {
		var result Commic
		var f *os.File
		var err error
		f, err = os.Open(r)
		if err != nil {
			fmt.Errorf("failed to read:%s", err)
			continue
		}

		dec := json.NewDecoder(f)
		if err := dec.Decode(&result); err != nil {
			fmt.Errorf("failed to read:%s", err)
			f.Close()
			continue
		}

		url := fmt.Sprintf("http://xkcd.com/%d", result.Num)
		result.Url = url

		if strings.Contains(result.Transcript, query) {
			fmt.Printf("-----\n%s\n-----\n%s\n", result.Url, result.Transcript)
		}

		f.Close()
	}
}

func listFiles(searchPath string) []string {
	fis, err := ioutil.ReadDir(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, fi := range fis {
		fullPath := filepath.Join(searchPath, fi.Name())
		files = append(files, fullPath)
	}

	return files
}
