package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("max(3,1,4)=")
	fmt.Println(max(3, 1, 4))
	fmt.Print("max()=")
	fmt.Println(max())

}

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("max:missing args")
	}

	val := math.MinInt64

	for _, v := range vals {
		if v > val {
			val = v
		}
	}
	return val, nil
}

func min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("min:missing args")
	}

	val := math.MaxInt64

	for _, v := range vals {
		if v < val {
			val = v
		}
	}
	return val, nil
}

func max2(val int, vals ...int) int {
	for _, v := range vals {
		if v > val {
			val = v
		}
	}
	return val
}

func min2(val int, vals ...int) int {
	for _, v := range vals {
		if v < val {
			val = v
		}
	}
	return val
}
