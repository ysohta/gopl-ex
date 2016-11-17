package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var temp = CelsiusFlag("temp", 20.0, "the temperature")
var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	fmt.Fprintln(out, *temp)
}
