package lin

import (
	"fmt"
)

type Num interface {
	IsNeg() bool
	Negate() Num
	Mul(Num) Num
	Add(Num) Num
	Sub(Num) Num
}

// A Term without Var is a constant
type Term struct {
	Coeff Num
	Var   string
}

func (t Term) Format(format string) string {
	sign := "+"
	if t.Coeff.IsNeg() {
		sign = "-"
	}
	return fmt.Sprintf(format, sign, t.Coeff, t.Var)
}

func (t Term) String() string {
	return t.Format("%s %v %s")
}

func (t Term) AsPoly() Poly {
	return Poly([]Term{t})
}
