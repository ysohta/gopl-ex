package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("query missing")
	}

	query := os.Args[1]

	result, err := search(query)
	if err != nil {
		log.Fatal(err)
	}

	if result.Response != "True" {
		log.Fatal("invalid response:", result.Error)
	}

	for _, m := range result.Search {
		if m.Poster == "N/A" {
			fmt.Printf("missing url. [%s]%s\n", m.ID, m.Title)
		} else {
			filename, err := downlaodPoster(*m)
			if err != nil {
				fmt.Printf("failed to download. %s\n", err)
			} else {
				fmt.Printf("download %s\n", filename)
			}
		}
	}
}
