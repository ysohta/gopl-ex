// This prints command name and the parameters with index.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(formatWithIndex(os.Args))
}

func formatWithIndex(strs []string) string {
	concat := ""
	for i, str := range strs {
		concat += fmt.Sprintf("[%d]%s\n", i, str)
	}
	return concat
}
