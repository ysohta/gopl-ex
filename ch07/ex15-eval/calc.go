package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/ysohta/gopl-ex/ch07/ex13-eval"
)

var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

func calc() {
	input := bufio.NewScanner(in)

	input.Scan()
	expr, err := eval.Parse(input.Text())
	if err != nil {
		fmt.Fprintf(out, "parse error: %s\n", err)
		return
	}

	vars := map[eval.Var]bool{}
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(out, "check error: %s\n", err)
		return
	}

	env := eval.Env{}

	// sort vars
	keys := make([]string, 0, len(vars))
	for v := range vars {
		keys = append(keys, string(v))
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := eval.Var(k)
		fmt.Fprintf(out, "%s: ", v)
		for input.Scan() {
			f, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(out, "parse error: %s\n", err)
				fmt.Fprintf(out, "%s: ", v)
				return
			}

			env[v] = f
			break
		}
	}
	got := fmt.Sprintf("%.6g", expr.Eval(env))
	fmt.Fprintln(out, got)

}
