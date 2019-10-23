package lin

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type stringTest struct {
	In  interface{}
	Out string
}

func newPolyTest(expect string, in ...string) stringTest {
	var p Poly
	for i := 0; i < len(in); i = i + 2 {
		val, _ := strconv.ParseInt(in[i], 10, 64)
		p = append(p, Term{IntCoeff(val), in[i+1]})
	}
	return stringTest{p, expect}
}

var stringTests = []stringTest{
	newPolyTest("2x - 5y", "2", "x", "-5", "y"),
	newPolyTest("- e + pi", "-1", "e", "1", "pi"),
	newPolyTest("- 5e + 4pi", "-5", "e", "4", "pi"),
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		str := test.In.(Poly).String()
		assertEquals(t, str, test.Out)
	}
}

var negateTests = []stringTest{
	newPolyTest("- 2x + 5y", "2", "x", "-5", "y"),
	newPolyTest("e - pi", "-1", "e", "1", "pi"),
	newPolyTest("5e - 4pi", "-5", "e", "4", "pi"),
}

func TestNegate(t *testing.T) {
	for _, test := range negateTests {
		str := test.In.(Poly).Negate().String()
		assertEquals(t, str, test.Out)
	}
}

var simplifyTests = []stringTest{
	newPolyTest("11x + 4y", "2", "x", "-5", "y", "9", "x", "9", "y"),
	newPolyTest("11x - 5y", "2", "x", "-5", "y", "9", "x"),
}

func TestSimplify(t *testing.T) {
	for _, test := range simplifyTests {
		str := test.In.(Poly).Simplify().String()
		assertEquals(t, str, test.Out)
	}
}

var composeTests = []stringTest{
	newPolyTest("4x - 10y + 18x + 18y - 5y + 18x - 45y + 81x + 81y + 9y",
		"2", "x", "-5", "y", "9", "x", "9", "y"),
	newPolyTest("4x - 10y + 18x + 5y + 18x - 45y + 81x", "-2", "x", "5", "y", "-9", "x"),
}

func TestCompose(t *testing.T) {
	for _, test := range composeTests {
		poly := test.In.(Poly)
		str := poly.Compose("x", poly).String()
		assertEquals(t, str, test.Out)
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

func assertEquals(t *testing.T, got, expect interface{}) {
	if got != expect {
		t.Errorf("got (%s), expected (%s)", got, expect)
	}
}

func randBool(r *rand.Rand) bool {
	return r.Int()%2 == 0
}
