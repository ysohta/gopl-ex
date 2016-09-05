package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	IssuesURL = "https://api.github.com/repos"
	jsonBody  = "application/json"
)

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
	Body   string `json:"body"`
}

func createIssue(issue Issue) (*Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues?access_token=%s", owner, repo, token)

	var b []byte
	var err error
	if b, err = json.Marshal(issue); err != nil {
		return nil, err
	}

	r := strings.NewReader(string(b))

	resp, err := http.Post(url, jsonBody, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ListIssues() ([]Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues?direction=asc&state=all&access_token=%s", owner, repo, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
