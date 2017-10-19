package lin

import (
	"fmt"
	"testing"
)

func TestReadPoly(t *testing.T) {
	var p Term
	p.Coeff = Frac64{10, 3, true}
	p.Var = "x"
	fmt.Println(p)
}
