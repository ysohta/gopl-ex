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

func GetIssue(number int) (*Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues/%d?access_token=%s", owner, repo, number, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result *Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func ListIssues() ([]Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues?direction=asc&state=all", owner, repo)

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

func editIssue(number int, issue Issue) (*Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues/%d?access_token=%s", owner, repo, number, token)

	body := strings.Replace(issue.Body, "\n", "\\n", -1)
	s := fmt.Sprintf("{\"title\":\"%s\", \"body\":\"%s\"}", issue.Title, body)

	r := strings.NewReader(s)

	resp, err := http.Post(url, jsonBody, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func closeIssue(number int) (*Issue, error) {
	url := fmt.Sprintf(IssuesURL+"/%s/%s/issues/%d?access_token=%s", owner, repo, number, token)

	r := strings.NewReader("{\"state\":\"closed\"}")

	resp, err := http.Post(url, jsonBody, r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
