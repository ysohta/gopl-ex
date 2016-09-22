package cmplxbgflt

import (
	"math/big"
	"testing"
)

func TestF(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
			&complexFloat{new(big.Float).SetFloat64(-1.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(1.0), new(big.Float).SetFloat64(0.0)},
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(-1.0), new(big.Float).SetFloat64(0.0)},
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(1.0)},
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(-1.0)},
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
		},
	}
	for _, test := range tests {
		got := f(test.x)
		if got.i.Cmp(test.want.i) != 0 {
			t.Errorf("got.i = %f, want.i = %f\n", got.i, test.want.i)
		}
		if got.r.Cmp(test.want.r) != 0 {
			t.Errorf("got.r = %f, want.r = %f\n", got.r, test.want.r)
		}
	}
}

func TestFd(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
			&complexFloat{new(big.Float).SetFloat64(0.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(1.0), new(big.Float).SetFloat64(0.0)},
			&complexFloat{new(big.Float).SetFloat64(4.0), new(big.Float).SetFloat64(0.0)},
		}, {
			&complexFloat{new(big.Float).SetFloat64(0.5), new(big.Float).SetFloat64(0.5)},
			&complexFloat{new(big.Float).SetFloat64(-1.0), new(big.Float).SetFloat64(1.0)},
		},
	}
	for _, test := range tests {
		got := fd(test.x)
		if got.i.Cmp(test.want.i) != 0 {
			t.Errorf("got.i = %f, want.i = %f\n", got.i, test.want.i)
		}
		if got.r.Cmp(test.want.r) != 0 {
			t.Errorf("got.r = %f, want.r = %f\n", got.r, test.want.r)
		}
	}
}
