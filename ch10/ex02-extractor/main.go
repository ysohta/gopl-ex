package main

import (
	"fmt"
	"os"

	"github.com/ysohta/gopl-ex/ch10/ex02-extractor/extractor"
	_ "github.com/ysohta/gopl-ex/ch10/ex02-extractor/zip" // register zip
	//_ "github.com/ysohta/gopl-ex/ch10/ex02-extractor/tar" // register tar
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("spcify archive file as argument")
		os.Exit(1)
	}

	dst := "."
	src := os.Args[1]
	if len(os.Args) > 2 {
		dst = os.Args[2]
	}

	err := extractor.Extract(dst, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "extract failure: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("success to extract file: %s to: %s\n", src, dst)
}
