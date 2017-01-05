package popcount

var (
	pc [256]byte
)

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the count of number of bits set to 1.
// ref) text 2.6.2
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}

// ShiftPopCount returns the count of number of bits set to 1.
// ref) ex 2.4
func ShiftPopCount(x uint64) int {
	var cnt int
	for i := 0; i < 64; i++ {
		cnt += int(x & 1)
		x >>= 1
	}
	return cnt
}

// PopCountClearMinBit returns the count of number of bits set to 1.
// ref) ex 2.5
func PopCountClearMinBit(x uint64) int {
	var cnt int
	for x != 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
