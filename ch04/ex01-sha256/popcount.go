package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the count of number of bits set to 1.
func PopCount(b []byte) int {
	var cnt int
	for i := 0; i < len(b); i++ {
		cnt += int(pc[b[i]])
	}
	return cnt
}
