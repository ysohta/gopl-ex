package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	latest := filter(result.Items, oneMonthAgo, now)
	fmt.Printf("\n[latest] %d issues\n", len(latest))
	for _, item := range latest {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	recent := filter(result.Items, oneYearAgo, oneMonthAgo)
	fmt.Printf("\n[recent] %d issues\n", len(recent))
	for _, item := range recent {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	old := filter(result.Items, time.Time{}, oneYearAgo)
	fmt.Printf("\n[old] %d issues\n", len(old))
	for _, item := range old {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func filter(items []*Issue, from, to time.Time) []*Issue {
	var filtered []*Issue

	for _, item := range items {
		if item.CreatedAt.After(from) && item.CreatedAt.Before(to) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
