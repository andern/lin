package lin

import (
	"fmt"
	"math/big"
)

// A Term without Var is a constant
type Term struct {
	Coeff big.Rat
	Var   string
}

func (t Term) Format(leftmost bool) string {
	sign := "+"
	if t.Coeff.Sign() == -1 {
		sign = "-"
	}

	val := new(big.Rat)
	val.Abs(&t.Coeff)
	return fmt.Sprintf(t.formatString(leftmost), sign, val.RatString(), t.Var)
}

// leftmost denotes whether it's the first (leftmost) term in a series of terms
func (t Term) formatString(leftmost bool) string {
	sign := "%s "
	coef := "%s"
	vari := "%s"

	if leftmost && t.Coeff.Sign() == 1 {
		sign = "%.0s"
	}

	if t.IsOne() && !t.IsConstant() {
		coef = "%.0s"
	}

	return sign + coef + vari
}

/*
func termFormat(idx int, t Term) string {
	sign := "%s "
	coef := "%v"
	vari := "%s"

	if idx == 0 && t.Coeff.Sign() == 1 {
		sign = "%.0s"
	}

	if t.IsOne() && !t.IsConstant() {
		coef = "%.0r"
	}

	return sign + coef + vari
} */

func (a Term) Add(b Term) (res Term, err error) {
	if a.Var != b.Var {
		err = fmt.Errorf("Term(Add): mismatched variable name")
	}
	res.Var = a.Var
	res.Coeff.Add(&a.Coeff, &b.Coeff)
	return
}

func (a Term) IsOne() bool {
	//	return a.Coeff.Num().Cmp(a.Coeff.Denom()) == 0
	return a.Coeff.Num().Abs(a.Coeff.Num()).Cmp(a.Coeff.Denom().Abs(a.Coeff.Denom())) == 0
	//	return a.Coeff.Denom().Cmp(big.NewInt(0)) != 0 && a.Coeff.Num() == a.Coeff.Denom()
}

func (a Term) IsConstant() bool {
	return a.Var == ""
}

func (t Term) String() string {
	return t.Format(false)
}

func (t Term) AsPoly() Poly {
	return Poly([]Term{t})
}
