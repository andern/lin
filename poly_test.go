package lin

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/andern/frac"
)

type testPoly struct {
	In  Poly
	Out string
}

var polyTests = []testPoly{
	{Poly([]Term{{frac.Frac64{2, 1, false}, "x"}, {frac.Frac64{5, 1, true}, "y"}}), "2/1 x - 5/1 y"},
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
			f := frac.Frac64{r.Uint64(), r.Uint64(), randBool(r)}
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
