package lin

import (
	"fmt"
)

type Num interface {
	Abs() Num
	Negate() Num
	Add(v Num) Num
	Mul(v Num) Num
	// Sign returns 1 if value is positive and -1 if value is negative. 0 otherwise
	Sign() int
	// IsSingle returns true if value is equal to 1 or -1
	IsSingle() bool
}

type Term struct {
	Coeff Num
	Var   string
}

func (t Term) IsConstant() bool {
	return t.Var == ""
}

func (t Term) Format(leftmost bool) string {
	sign := "+"
	if t.Coeff.Sign() == -1 {
		sign = "-"
	}

	return fmt.Sprintf(t.formatString(leftmost), sign, t.Coeff.Abs(), t.Var)
}

func (t Term) String() string {
	return t.Format(false)
}

func (t Term) AsPoly() Poly {
	return Poly([]Term{t})
}

// leftmost denotes whether it's the first (leftmost) term in a series of terms
func (t Term) formatString(leftmost bool) string {
	sign := "%s "
	coef := "%s"
	vari := "%s"

	if leftmost && t.Coeff.Sign() == 1 {
		sign = "%.0s"
	}

	if t.Coeff.IsSingle() && !t.IsConstant() {
		coef = "%.0s"
	}

	return sign + coef + vari
}
