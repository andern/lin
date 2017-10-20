package lin

import (
	"sort"
)

// A Poly represents a linear polynomial
type Poly []Term

// Poly implements sort.Interface
func (p Poly) Len() int           { return len(p) }
func (p Poly) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Poly) Less(i, j int) bool { return p[i].Var < p[j].Var }

// Add adds a polynomial to another polynomial without simplifying the result
func (p Poly) Add(p2 Poly) Poly {
	return append(p, p2...)
}

// Sub subtracts a polynomial to another
// polynomial without simplifying the result
func (p Poly) Sub(p2 Poly) Poly {
	return p.Add(p2.Negate())
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

func (p Poly) Negate() Poly {
	for i, term := range p {
		p[i] = term.Negate()
	}
	return p
}
