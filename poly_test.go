package lin

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type stringTest struct {
	In  Poly
	Out string
}

func newStringTest(expect string, in ...string) stringTest {
	var p Poly
	for i := 0; i < len(in); i = i + 2 {
		val, _ := strconv.ParseInt(in[i], 10, 64)
		p = append(p, Term{IntCoeff(val), in[i+1]})
	}
	return stringTest{p, expect}
}

var stringTests = []stringTest{
	newStringTest("2x - 5y", "2", "x", "-5", "y"),
	newStringTest("- e + pi", "-1", "e", "1", "pi"),
	newStringTest("- 5e + 4pi", "-5", "e", "4", "pi"),
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
	newStringTest("- 2x + 5y", "2", "x", "-5", "y"),
	newStringTest("e - pi", "-1", "e", "1", "pi"),
	newStringTest("5e - 4pi", "-5", "e", "4", "pi"),
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
	newStringTest("11x + 4y", "2", "x", "-5", "y", "9", "x", "9", "y"),
	newStringTest("11x - 5y", "2", "x", "-5", "y", "9", "x"),
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

var composeTests = []stringTest{
	newStringTest("4x - 10y + 18x + 18y - 5y + 18x - 45y + 81x + 81y + 9y",
		"2", "x", "-5", "y", "9", "x", "9", "y"),
	newStringTest("4x - 10y + 18x + 5y + 18x - 45y + 81x", "-2", "x", "5", "y", "-9", "x"),
}

func TestCompose(t *testing.T) {
	for _, test := range composeTests {
		str := test.In.Compose("x", test.In).String()
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
			f := r.Int63()
			term := Term{IntCoeff(f), "x"}
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
