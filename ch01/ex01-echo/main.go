// This prints command name and the parameters.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	// => "./ex01-echo param1 param2"
}
