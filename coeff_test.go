package lin

import (
	"math/big"
	"testing"
)

func TestIsSingle(t *testing.T) {
	check := func(t *testing.T, coeff Coeff, expected bool) {
		t.Helper()
		assertEquals(t, coeff.IsSingle(), expected)
	}

	t.Run("IntCoeff", func(t *testing.T) {
		check(t, IntCoeff(-2), false)
		check(t, IntCoeff(-1), true)
		check(t, IntCoeff(0), false)
		check(t, IntCoeff(1), true)
		check(t, IntCoeff(2), false)
	})

	t.Run("FloatCoeff", func(t *testing.T) {
		check(t, FloatCoeff(-123789), false)
		check(t, FloatCoeff(-2), false)
		check(t, FloatCoeff(-1.1), false)
		check(t, FloatCoeff(-1.0001), false)
		check(t, FloatCoeff(-0.9999), false)
		check(t, FloatCoeff(0), false)
		check(t, FloatCoeff(0.9999), false)
		check(t, FloatCoeff(1), true)
		check(t, FloatCoeff(1.0001), false)
		check(t, FloatCoeff(2), false)
		check(t, FloatCoeff(98243), false)
	})

	t.Run("RatCoeff", func(t *testing.T) {
		check(t, RatCoeff{big.NewRat(-4123708, 2340789)}, false)
		check(t, RatCoeff{big.NewRat(-4123708, -2340789)}, false)
		check(t, RatCoeff{big.NewRat(-4123708, 4123708)}, true)
		check(t, RatCoeff{big.NewRat(-2, 1)}, false)
		check(t, RatCoeff{big.NewRat(-2, 2)}, true)
		check(t, RatCoeff{big.NewRat(-1, 1)}, true)
		check(t, RatCoeff{big.NewRat(0, -1)}, false)
		check(t, RatCoeff{big.NewRat(0, -2)}, false)
		check(t, RatCoeff{big.NewRat(0, 1)}, false)
		check(t, RatCoeff{big.NewRat(0, 2)}, false)
		check(t, RatCoeff{big.NewRat(1, -1)}, true)
		check(t, RatCoeff{big.NewRat(1, 1)}, true)
		check(t, RatCoeff{big.NewRat(2, 1)}, false)
		check(t, RatCoeff{big.NewRat(2, -1)}, false)
		check(t, RatCoeff{big.NewRat(2, 2)}, true)
		check(t, RatCoeff{big.NewRat(4123708, 2340789)}, false)
		check(t, RatCoeff{big.NewRat(4123708, -2340789)}, false)
		check(t, RatCoeff{big.NewRat(4123708, 4123708)}, true)
	})
}

func TestSign(t *testing.T) {
	check := func(t *testing.T, coeff Coeff, expected int) {
		t.Helper()
		assertEquals(t, coeff.Sign(), expected)
	}

	t.Run("IntCoeff", func(t *testing.T) {
		check(t, IntCoeff(-2), -1)
		check(t, IntCoeff(-1), -1)
		check(t, IntCoeff(0), 0)
		check(t, IntCoeff(1), 1)
		check(t, IntCoeff(2), 1)
	})

	t.Run("FloatCoeff", func(t *testing.T) {
		check(t, FloatCoeff(-2), -1)
		check(t, FloatCoeff(-1), -1)
		check(t, FloatCoeff(-0.0001), -1)
		check(t, FloatCoeff(0), 0)
		check(t, FloatCoeff(0.0001), 1)
		check(t, FloatCoeff(1), 1)
		check(t, FloatCoeff(2), 1)
	})
}

func TestAbs(t *testing.T) {
	check := func(t *testing.T, coeff, expected Coeff) {
		t.Helper()
		assertEquals(t, coeff.Abs(), expected)
	}

	t.Run("IntCoeff", func(t *testing.T) {
		check(t, IntCoeff(-2), IntCoeff(2))
		check(t, IntCoeff(-1), IntCoeff(1))
		check(t, IntCoeff(0), IntCoeff(0))
		check(t, IntCoeff(1), IntCoeff(1))
		check(t, IntCoeff(2), IntCoeff(2))
	})

	t.Run("FloatCoeff", func(t *testing.T) {
		check(t, FloatCoeff(-2), FloatCoeff(2))
		check(t, FloatCoeff(-1), FloatCoeff(1))
		check(t, FloatCoeff(-0.0001), FloatCoeff(0.0001))
		check(t, FloatCoeff(0), FloatCoeff(0))
		check(t, FloatCoeff(0.0001), FloatCoeff(0.0001))
		check(t, FloatCoeff(1), FloatCoeff(1))
		check(t, FloatCoeff(2), FloatCoeff(2))
	})
}
