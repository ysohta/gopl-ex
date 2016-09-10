package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const PosterURL = "https://omdbapi.com/"

type Result struct {
	Response string   `json:"Response"`
	Error    string   `json:"Error"`
	Search   []*Movie `json:"Search"`
}

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
	ID     string `json:"imdbID"`
}

func search(query string) (*Result, error) {
	resp, err := http.Get(PosterURL + "?s=" + query)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func downlaodPoster(m Movie) (string, error) {
	resp, err := http.Get(m.Poster)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", fmt.Errorf("download image fail: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	title := strings.Replace(m.Title, " ", "_", -1)
	filename := fmt.Sprintf("%s_%s.jpg", m.ID, title)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.Write(body)
	return filename, nil
}
