package lin

type Equality string

const (
	GreaterThan        Equality = ">"
	GreaterThanOrEqual          = ">="
	Equal                       = "="
	NotEqual                    = "!="
	LessThan                    = "<"
	LessThanOrEqual             = "<="
)

type Expression struct {
	LHS Poly
	RHS Poly
}

type Equation struct {
	Expression
}

type Inequality struct {
	Expression
	Equality
}
