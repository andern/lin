package lin

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/andern/frac"
)

type Frac frac.Frac64

// Implement Num
func (f Frac) IsNeg() bool   { return f.Neg }
func (f Frac) IsZero() bool  { return f.Num == 0 }
func (f Frac) IsOne() bool   { return f.Num == f.Den }
func (f Frac) Negate() Num   { return f.Negate() }
func (f Frac) Mul(v Num) Num { return f.Mul(v) }
func (f Frac) Add(v Num) Num { return f.Add(v) }

func (f Frac) String() string {
	if f.Den == 1 {
		return fmt.Sprintf("%v", f.Num)
	}
	return fmt.Sprintf("%v/%v", f.Num, f.Den)
}

type testPoly struct {
	In  Poly
	Out string
}

var polyTests = []testPoly{
	{Poly([]Term{{Frac{2, 3, false}, "x"}, {Frac{5, 3, true}, "y"}}), "2/3x - 5/3y"},
	{Poly([]Term{{Frac{2, 2, true}, "e"}, {Frac{4, 4, false}, "pi"}}), "- e + pi"},
	{Poly([]Term{{Frac{5, 1, true}, "e"}, {Frac{4, 1, false}, "pi"}}), "- 5e + 4pi"},
}

func TestString(t *testing.T) {
	for _, test := range polyTests {
		str := test.In.String()
		err := checkOutput(str, test.Out)
		if err != nil {
			t.Error(err)
		}
	}
}

func BenchmarkString(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < b.N; n++ {
		var p Poly
		for m := 0; m < r.Int()%10; m++ {
			f := Frac{r.Uint64(), r.Uint64(), randBool(r)}
			term := Term{f, "x"}
			p = append(p, term)
		}
		p.String()
	}
}

func checkOutput(got, expect string) error {
	if got == expect {
		return nil
	}
	return errors.New(fmt.Sprintf("got (%s), expected (%s)", got, expect))
}

func randBool(r *rand.Rand) bool {
	return r.Int()%2 == 0
}
