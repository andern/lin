package lin

import (
	"bytes"
	"github.com/andern/frac"
)

type Term struct {
	Coeff frac.Frac64
	Var   string
}

func FracString(f frac.Frac64) string {
	var buffer bytes.Buffer
	if f.Neg {
		buffer.WriteString("-")
	}
	buffer.WriteString(string(f.Num))
	buffer.WriteString("/")
	buffer.WriteString(string(f.Den))
	return buffer.String()
}

func (t *Term) String() string {
	return FracString(t.Coeff) + t.Var
}
