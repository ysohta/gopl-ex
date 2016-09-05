package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cmd, owner, repo, token, number, editor string
)

func init() {
	flag.StringVar(&cmd, "cmd", "list", "command for the issues")
	flag.StringVar(&owner, "owner", "", "repository owner")
	flag.StringVar(&repo, "repo", "", "repository name")
	flag.StringVar(&token, "token", "", "OAUTH token")
	flag.StringVar(&number, "number", "", "issue number")
	flag.StringVar(&editor, "editor", "nano", "text editor program")

	flag.Parse()
}

func main() {
	switch cmd {
	case "create":
		fmt.Print("title:")
		title := readLine()
		fmt.Print("body:")
		body := readLine()
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
	fmt.Println(issue.Body)
}
