package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash string

func init() {
	flag.StringVar(&hash, "hash", "sha256", "hash algorithm(sha256, sha384 or sha512)")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch hash {
		case "sha384":
			c := sha512.Sum384(scanner.Bytes())
			fmt.Printf("%x\n", c)
		case "sha512":
			c := sha512.Sum512(scanner.Bytes())
			fmt.Printf("%x\n", c)
		default:
			c := sha256.Sum256(scanner.Bytes())
			fmt.Printf("%x\n", c)
		}
	}
}
