package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing args\n")
		os.Exit(1)
	}
	arg := os.Args[1]
	val, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("cannot convert:%v\n", err)
		os.Exit(1)
	}

	catchPanic(val)
}

type bailout struct{ ret int }

func catchPanic(val int) {
	defer func() {
		p := recover()
		if p != nil {
			b, ok := p.(bailout)
			if ok {
				fmt.Printf("ret=%d\n", b.ret)
			}
		}
	}()

	noReturn(val)
}

func noReturn(val int) {
	panic(bailout{val})
}
