package lin

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

type stringTest struct {
	In  Poly
	Out string
}

var stringTests = []stringTest{
	{Poly([]Term{{*big.NewRat(2, 3), "x"}, {*big.NewRat(-5, 3), "y"}}), "2/3x - 5/3y"},
	{Poly([]Term{{*big.NewRat(-2, 2), "e"}, {*big.NewRat(4, 4), "pi"}}), "- e + pi"},
	{Poly([]Term{{*big.NewRat(-5, 1), "e"}, {*big.NewRat(4, 1), "pi"}}), "- 5e + 4pi"},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		str := test.In.String()
		err := checkOutput(str, test.Out)
		if err != nil {
			t.Error(err)
		}
	}
}

var negateTests = []stringTest{
	{Poly([]Term{{*big.NewRat(2, 3), "x"}, {*big.NewRat(-5, 3), "y"}}), "- 2/3x + 5/3y"},
	{Poly([]Term{{*big.NewRat(-2, 2), "e"}, {*big.NewRat(4, 4), "pi"}}), "e - pi"},
	{Poly([]Term{{*big.NewRat(-5, 1), "e"}, {*big.NewRat(4, 1), "pi"}}), "5e - 4pi"},
}

func TestNegate(t *testing.T) {
	for _, test := range negateTests {
		str := test.In.Negate().String()
		err := checkOutput(str, test.Out)
		if err != nil {
			t.Error(err)
		}
	}
}

var simplifyTests = []stringTest{
	{Poly([]Term{
		{*big.NewRat(2, 3), "x"},
		{*big.NewRat(-5, 3), "y"},
		{*big.NewRat(3, 1), "x"},
		{*big.NewRat(3, 1), "y"},
	}), "11/3x + 4/3y"},

	{Poly([]Term{
		{*big.NewRat(2, 3), "x"},
		{*big.NewRat(-5, 3), "y"},
		{*big.NewRat(3, 1), "x"},
	}), "11/3x - 5/3y"},
}

func TestSimplify(t *testing.T) {
	for _, test := range simplifyTests {
		str := test.In.Simplify().String()
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
			f := big.NewRat(r.Int63(), r.Int63())
			term := Term{*f, "x"}
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
