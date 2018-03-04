package lin

import (
	"bytes"
	"sort"
)

// A Poly represents a linear polynomial
type Poly []Term

// Poly implements sort.Interface
func (p Poly) Len() int           { return len(p) }
func (p Poly) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Poly) Less(i, j int) bool { return p[i].Var < p[j].Var }

// Add adds a polynomial to another polynomial without simplifying the result
func (p Poly) Add(poly Poly) Poly {
	return append(p, poly...)
}

// Sub subtracts a polynomial to another
// polynomial without simplifying the result
func (p Poly) Sub(poly Poly) Poly {
	return p.Add(poly.Negate())
}

// Simplify simplifies a polynomial by combining
// like terms and reordering terms alphabetically
func (p Poly) Simplify() (res Poly) {
	sort.Sort(p)

	for i := 0; i < len(p)-1; i++ {
		term := Term{p[i].Coeff, p[i].Var}
		for ; i+1 < len(p) && term.Var == p[i+1].Var; i++ {
			term.Coeff = term.Coeff.Add(p[i+1].Coeff)
		}
		res = append(res, term)
	}
	return
}

// Negate negates all terms in the polynomial
func (p Poly) Negate() Poly {
	for i, _ := range p {
		p[i].Coeff = p[i].Coeff.Negate()
	}
	return p
}

// Compose substitutes a variable in p with the given poly
func (p Poly) Compose(variable string, poly Poly) (res Poly) {
	for _, t := range p {
		if t.Var != variable {
			res = append(res, t)
			continue
		}
		for _, term := range poly {
			term.Coeff = term.Coeff.Mul(t.Coeff)
			res = append(res, term)
		}
	}
	return
}

func (p Poly) String() string {
	var b bytes.Buffer

	for i, term := range p {
		fmt := termFormat(i, term)
		b.WriteString(term.Format(fmt))

		if i < len(p)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func termFormat(idx int, t Term) string {
	sign := "%s"
	coef := "%v"
	vari := "%s"

	if idx == 0 && !t.Coeff.IsNeg() {
		sign = "%.0s"
	}

	if t.Coeff.IsOne() && t.Var != "" {
		coef = "%.0v"
	}

	if idx == 0 {
		return sign + coef + vari
	}
	return sign + " " + coef + vari
}
