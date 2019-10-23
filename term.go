package lin

import (
	"fmt"
)

type Term struct {
	Coeff Coeff
	Var   string
}

func (t Term) IsConstant() bool {
	return t.Coeff.Sign() == 0 || t.Var == ""
}

func (t Term) Format(leftmost bool) string {
	if t.Coeff.Sign() == 0 {
		return ""
	}

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
