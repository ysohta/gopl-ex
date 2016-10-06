package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	courses, acyclic := topoSort(prereqs)
	if acyclic {
		fmt.Println("acyclic found")
	}
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	visited := make(map[string]bool)
	var visitAll func(items []string, visited map[string]bool) bool

	visitAll = func(items []string, visited map[string]bool) (acyclic bool) {
		for _, item := range items {
			if visited[item] {
				return true
			}
			if !seen[item] {
				seen[item] = true

				visited[item] = true
				acyclic = visitAll(m[item], visited)
				visited[item] = false

				order = append(order, item)

			}
		}
		return acyclic
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	acyclic := visitAll(keys, visited)
	if acyclic {
		order = nil
	}
	return order, acyclic
}
