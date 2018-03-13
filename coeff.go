package lin

import (
	"math/big"
	"strconv"
)

type Coeff interface {
	Abs() Coeff
	Negate() Coeff
	Add(v Coeff) Coeff
	Mul(v Coeff) Coeff
	// Sign returns 1 if value is positive and -1 if value is negative. 0 otherwise
	Sign() int
	// IsSingle returns true if value is equal to 1 or -1
	IsSingle() bool
}

// Convenience types for common cases
type IntCoeff int64

func (a IntCoeff) IsSingle() bool    { return a == 1 || a == -1 }
func (a IntCoeff) Add(b Coeff) Coeff { return a + b.(IntCoeff) }
func (a IntCoeff) Mul(b Coeff) Coeff { return a * b.(IntCoeff) }
func (a IntCoeff) Negate() Coeff     { return -a }
func (a IntCoeff) String() string    { return strconv.FormatInt(int64(a), 10) }
func (a IntCoeff) Sign() int {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	return 0
}
func (a IntCoeff) Abs() Coeff {
	if a < 0 {
		return -a
	}
	return a
}

type FloatCoeff float64

func (a FloatCoeff) IsSingle() bool    { return a == 1 || a == -1 }
func (a FloatCoeff) Add(b Coeff) Coeff { return a + b.(FloatCoeff) }
func (a FloatCoeff) Mul(b Coeff) Coeff { return a * b.(FloatCoeff) }
func (a FloatCoeff) Negate() Coeff     { return -a }
func (a FloatCoeff) String() string    { return strconv.FormatFloat(float64(a), 'E', -1, 64) }
func (a FloatCoeff) Sign() int {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	return 0
}
func (a FloatCoeff) Abs() Coeff {
	if a < 0 {
		return -a
	}
	return a
}

type RatCoeff struct {
	*big.Rat
}

func (a RatCoeff) IsSingle() bool {
	num := new(big.Int)
	denom := new(big.Int)
	num.Abs(a.Num())
	denom.Abs(a.Denom())

	return num.Cmp(denom) == 0
}
func (a RatCoeff) Add(b Coeff) Coeff {
	c := new(big.Rat)
	c.Set(a.Rat)
	c.Add(c, b.(RatCoeff).Rat)
	return RatCoeff{c}
}
func (a RatCoeff) Mul(b Coeff) Coeff {
	c := new(big.Rat)
	c.Set(a.Rat)
	c.Mul(c, b.(RatCoeff).Rat)
	return RatCoeff{c}
}
func (a RatCoeff) Negate() Coeff {
	c := new(big.Rat)
	c.Set(a.Rat)
	c.Neg(c)
	return RatCoeff{c}
}
func (a RatCoeff) String() string { return a.Rat.RatString() }
func (a RatCoeff) Abs() Coeff {
	c := new(big.Rat)
	c.Set(a.Rat)
	c.Abs(c)
	return RatCoeff{c}
}
