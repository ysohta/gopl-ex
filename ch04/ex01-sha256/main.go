package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	diff := diffBitSha256("x", "X")
	fmt.Printf("%d\n", diff)
	// =>"125"
}

func diffBitSha256(data1, data2 string) int {
	c1 := sha256.Sum256([]byte(data1))
	c2 := sha256.Sum256([]byte(data2))

	diff := xor(c1, c2)

	return PopCount(diff[0:len(diff)])
}

func xor(b1, b2 [32]byte) [32]byte {
	b := [32]byte{}
	for i := range b1 {
		b[i] = b1[i] ^ b2[i]
	}
	return b
}
