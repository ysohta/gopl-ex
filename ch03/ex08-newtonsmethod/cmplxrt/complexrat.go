package cmplxrt

import "math/big"

type complexRat struct {
	r *big.Rat
	i *big.Rat
}

func newComplexRat() *complexRat {
	cr := new(complexRat)
	cr.r = new(big.Rat)
	cr.i = new(big.Rat)
	return cr
}

func add(x, y *complexRat) *complexRat {
	z := newComplexRat()
	z.r.Add(x.r, y.r)
	z.i.Add(x.i, y.i)
	return z
}

func neg(x *complexRat) *complexRat {
	z := newComplexRat()
	z.r.Neg(x.r)
	z.i.Neg(x.i)
	return z
}

func mul(x, y *complexRat) *complexRat {
	t1 := new(big.Rat)
	t2 := new(big.Rat)
	z := newComplexRat()
	t1.Mul(x.r, y.r)
	t2.Mul(x.i, y.i)
	z.r.Sub(t1, t2)

	t1.Mul(x.r, y.i)
	t2.Mul(x.i, y.r)
	z.i.Add(t1, t2)
	return z
}

func quo(x, y *complexRat) *complexRat {
	z := newComplexRat()
	denominator := new(big.Rat)
	t := new(big.Rat)
	t.Mul(y.r, y.r)
	denominator.Mul(y.i, y.i)
	denominator.Add(denominator, t)

	if denominator.Cmp(zero) == 0 {
		return newComplexRat()
	}

	ac := new(big.Rat)
	bd := new(big.Rat)
	ac.Mul(x.r, y.r)
	bd.Mul(x.i, y.i)

	bc := new(big.Rat)
	ad := new(big.Rat)
	bc.Mul(x.i, y.r)
	ad.Mul(x.r, y.i)

	z.r.Add(ac, bd)
	z.r.Quo(z.r, denominator)

	z.i.Add(bc, ad.Neg(ad))
	z.i.Quo(z.i, denominator)

	return z
}

func abs(x *complexRat) *big.Rat {
	r := new(big.Rat)
	i := new(big.Rat)
	r.Set(x.r)
	i.Set(x.i)

	r.Mul(r, x.r) // r^2
	i.Mul(i, x.i) // i^2

	r.Add(r, i) // r^2 + i^2

	return sqrtFloat(r) // sqrt(r^2 + i^2)
}

func sqrtFloat(x *big.Rat) *big.Rat {
	t1 := new(big.Rat)
	t2 := new(big.Rat)
	t1.Set(x)

	// Iterate.
	// x{n} = (x{n-1}+x{0}/x{n-1}) / 2
	for i := 0; i <= 4; i++ {
		if t1.Cmp(zero) == 0 {
			return t1
		}
		t2.Quo(x, t1)
		t2.Add(t2, t1)
		t1.Mul(half, t2)
	}

	return t1
}
