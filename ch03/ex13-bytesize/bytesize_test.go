package bytesize

import (
	"testing"
)

func TestBytesize(t *testing.T) {
	tests := []struct {
		b    uint64
		want uint64
	}{
		{KiB, 1024},
		{MiB, 1048576},
		{GiB, 1073741824},
		{TiB, 1099511627776},
		{PiB, 1125899906842624},
		{EiB, 1152921504606846976},
		{ZiB / EiB, 1024}, // avoid overflow
		{YiB / ZiB, 1024}, // avoid overflow
	}

	for _, test := range tests {
		if test.b != test.want {
			t.Errorf("Actual: %d\tExpected: %d", test.b, test.want)
		}
	}
}
