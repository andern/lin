package lin

import (
	"fmt"

	"github.com/andern/frac"
)

// A Term without Var is a constant
type Term struct {
	Coeff frac.Frac64
	Var   string
}

func (t Term) Format(format string) string {
	sign := "+"
	if t.Coeff.Neg {
		sign = "-"
	}
	return fmt.Sprintf(format, sign, t.Coeff.Num, t.Coeff.Den, t.Var)
}

func (t Term) String() string {
	return t.Format("%s %v/%v %s")
}

func (t Term) AsPoly() Poly {
	return Poly([]Term{t})
}
