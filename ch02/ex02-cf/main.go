// ex02-cf prints values converted by general unit converter.
package main

import (
	"bufio"
	"fmt"
	"github.com/ysohta/gopl-ex/ch02/ex01-tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		// convert from stdard input
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			conv(input.Text())
		}
	} else {
		// convert from arguments
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	}
}

func conv(val string) {
	n, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(convertTemps(n))
	fmt.Println(convertDistances(n))
	fmt.Println(convertWeights(n))
}

func convertTemps(n float64) string {
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	return fmt.Sprintf("%s = %s, %s = %s", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func convertDistances(n float64) string {
	ft := Feet(n)
	m := Meter(n)
	return fmt.Sprintf("%s = %s, %s = %s", ft, FtToM(ft), m, MToFt(m))
}

func convertWeights(n float64) string {
	lb := Pound(n)
	kg := Kilogram(n)
	return fmt.Sprintf("%s = %s, %s = %s", lb, LbToKg(lb), kg, KgToLb(kg))
}
