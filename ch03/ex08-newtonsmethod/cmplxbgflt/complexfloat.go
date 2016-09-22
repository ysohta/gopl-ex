package cmplxbgflt

import (
	"math"
	"math/big"
)

type complexFloat struct {
	r *big.Float
	i *big.Float
}

var steps = int(math.Log2(float64(prec)))

func newComplexFloat() *complexFloat {
	cf := new(complexFloat)
	cf.r = new(big.Float).SetPrec(prec)
	cf.i = new(big.Float).SetPrec(prec)
	return cf
}

func add(x, y *complexFloat) *complexFloat {
	z := newComplexFloat()
	z.r.Add(x.r, y.r)
	z.i.Add(x.i, y.i)
	return z
}

func neg(x *complexFloat) *complexFloat {
	z := newComplexFloat()
	z.r.Neg(x.r)
	z.i.Neg(x.i)
	return z
}

func mul(x, y *complexFloat) *complexFloat {
	t1 := new(big.Float).SetPrec(prec)
	t2 := new(big.Float).SetPrec(prec)
	z := newComplexFloat()
	t1.Mul(x.r, y.r)
	t2.Mul(x.i, y.i)
	t2.Neg(t2)
	z.r.Add(t1, t2)

	t1.Mul(x.r, y.i)
	t2.Mul(x.i, y.r)
	z.i.Add(t1, t2)
	return z
}

func quo(x, y *complexFloat) *complexFloat {
	z := newComplexFloat()
	denominator := new(big.Float).SetPrec(prec)
	c2 := new(big.Float).SetPrec(prec)
	d2 := new(big.Float).SetPrec(prec)
	c2.Mul(y.r, y.r)
	d2.Mul(y.i, y.i)
	denominator.Add(c2, d2)

	if denominator.Cmp(zero) == 0 || denominator.IsInf() {
		return newComplexFloat()
	}

	ac := new(big.Float).SetPrec(prec)
	bd := new(big.Float).SetPrec(prec)
	ac.Mul(x.r, y.r)
	bd.Mul(x.i, y.i)

	bc := new(big.Float).SetPrec(prec)
	ad := new(big.Float).SetPrec(prec)
	bc.Mul(x.i, y.r)
	ad.Mul(x.r, y.i)

	z.r.Add(ac, bd)
	z.r.Quo(z.r, denominator)

	z.i.Add(bc, ad.Neg(ad))
	z.i.Quo(z.i, denominator)

	return z
}

func abs(x *complexFloat) *big.Float {
	r := new(big.Float).SetPrec(prec)
	i := new(big.Float).SetPrec(prec)
	r.Copy(x.r)
	i.Copy(x.i)

	r.Mul(r, x.r) // r^2
	i.Mul(i, x.i) // i^2

	r.Add(r, i) // r^2 + i^2

	return sqrtFloat(r) // sqrt(r^2 + i^2)
}

func sqrtFloat(x *big.Float) *big.Float {
	t1 := new(big.Float).SetPrec(prec)
	t2 := new(big.Float).SetPrec(prec)
	t1.Copy(x)

	// Iterate.
	// x{n} = (x{n-1}+x{0}/x{n-1}) / 2
	for i := 0; i <= steps; i++ {
		if t1.Cmp(zero) == 0 || t1.IsInf() {
			return t1
		}
		t2.Quo(x, t1)
		t2.Add(t2, t1)
		t1.Mul(half, t2)
	}

	return t1
}
