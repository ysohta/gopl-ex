package popcount

// PopCountClearMinBit returns the count of number of bits set to 1.
func PopCountClearMinBit(x uint64) int {
	var cnt int
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
