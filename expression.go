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
