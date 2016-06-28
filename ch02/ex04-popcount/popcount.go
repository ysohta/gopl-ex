package popcount

// ShiftPopCount returns the count of number of bits set to 1.
func ShiftPopCount(x uint64) int {
	var cnt int
	for i := 0; i < 64; i++ {
		cnt += int(x & 1)
		x >>= 1
	}
	return cnt
}
