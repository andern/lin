package lin

type Equality string

const (
	Equal              Equality = "="
	NotEqual                    = "!="
	GreaterThan                 = ">"
	GreaterThanOrEqual          = ">="
	LessThan                    = "<"
	LessThanOrEqual             = "<="
)

type Expression struct {
	LHS Poly
	Equality
	RHS Poly
}

// Compose substitutes a variable in e with the given poly
func (e *Expression) Compose(variable string, poly Poly) (res Expression) {
	res.LHS = e.LHS.Compose(variable, poly)
	res.RHS = e.RHS.Compose(variable, poly)
	res.Equality = e.Equality
	return
}

// Simplify simplifies an expression by subtracting on both sides - the RHS
// coefficients to variables found on the LHS - and then simplifying both sides
func (e *Expression) Simplify() Expression {
	//	res.LHS = e.LHS.Sub(e.RHS.Negate()).Simplify()

	lhs := e.LHS.Simplify()
	rhs := e.RHS.Simplify()

	for _, l := range e.LHS.Simplify() {
		for _, r := range e.RHS.Simplify() {
			if l.Var == r.Var {
				lhs.Sub(NewPoly(r))
				rhs.Sub(NewPoly(r))
			}
		}
	}

	return Expression{lhs.Simplify(), e.Equality, rhs.Simplify()}
}

func (e *Expression) SimplifyToLHS() Expression {
	return Expression{
		e.LHS.Sub(e.RHS).Simplify(),
		e.Equality,
		NewPoly(Term{}),
	}
}
