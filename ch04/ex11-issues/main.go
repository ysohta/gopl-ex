package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	cmd, owner, repo, token, editor string
)

func init() {
	flag.StringVar(&cmd, "cmd", "list", "command for the issues [create, list, edit, close]")
	flag.StringVar(&owner, "owner", "", "github epository owner")
	flag.StringVar(&repo, "repo", "", "github epository name")
	flag.StringVar(&token, "token", "", "OAUTH token")
	flag.StringVar(&editor, "editor", "vi", "text editor program")

	flag.Parse()
}

func main() {
	switch cmd {
	case "create":
		var title, body string
		var err error
		fmt.Print("title:")
		title = readLine()
		if err != nil {
			log.Fatal(err)
		}
		body, err = execute(editor, "Leave a comment")
		issue, err := createIssue(Issue{Title: title, Body: body})
		if err != nil {
			log.Fatal(err)
		}
		printIssue(*issue)

	case "list":
		issues, err := ListIssues()
		if err != nil {
			log.Fatal(err)
		}
		printIssues(issues)

	case "edit":
		fmt.Print("number:")
		s := readLine()

		var i int
		var err error
		var issue *Issue
		if i, err = strconv.Atoi(s); err != nil {
			log.Fatal(err)
		}

		issue, err = GetIssue(i)
		if err != nil {
			log.Fatal(err)
		}
		printIssue(*issue)

		var title, body string
		fmt.Print("title:")
		title = readLine()
		if err != nil {
			log.Fatal(err)
		}
		body, err = execute(editor, issue.Body)
		issue, err = editIssue(i, Issue{Title: title, Body: body})
		if err != nil {
			log.Fatal(err)
		}
		printIssue(*issue)

	case "close":
		fmt.Print("number:")
		s := readLine()
		var i int
		var err error
		if i, err = strconv.Atoi(s); err != nil {
			log.Fatal(err)
		}
		issue, err := closeIssue(i)
		if err != nil {
			log.Fatal(err)
		}
		printIssue(*issue)
	}
}

func readLine() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func printIssues(issues []Issue) {
	for _, issue := range issues {
		printIssue(issue)
	}
}

func printIssue(issue Issue) {
	fmt.Println("-----")
	fmt.Printf("#%d[%s]%s\n", issue.Number, issue.State, issue.Title)
	fmt.Printf("%s\n\n", issue.Body)
}
