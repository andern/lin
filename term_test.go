package lin

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/andern/frac"
)

func TestWhatever(t *testing.T) {
	a := Term{frac.Frac64{10, 3, true}, "a"}.AsPoly()
	b := Term{frac.Frac64{7, 13, false}, "b"}.AsPoly()
	c := Term{frac.Frac64{13, 5, false}, "a"}.AsPoly()

	p := a.Add(b).Add(b).Add(a).Add(a).Add(c).Add(c).Add(c).Add(c).Add(c)
	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
	fmt.Println(p.Simplify())
	fmt.Println(p.Simplify().Negate())
}

func BenchmarkFracFmt(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < b.N; n++ {
		f := frac.Frac64{r.Uint64(), r.Uint64(), randBool(r)}

		fmt.Sprintf("%v %v/%v %v", "-", f.Num, f.Den, "x")
	}
}

func randBool(r *rand.Rand) bool {
	return r.Int()%2 == 0
}
