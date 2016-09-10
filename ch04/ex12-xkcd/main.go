package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("query missing")
		os.Exit(1)
	}

	query := os.Args[1]
	fmt.Println("search query:", query)

	search(query)
}
