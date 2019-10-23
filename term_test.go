package lin

import (
	"testing"
)

func TestIsConstant(t *testing.T) {
	assertEquals(t, Term{IntCoeff(0), "x"}.IsConstant(), true)
	assertEquals(t, Term{IntCoeff(1), "x"}.IsConstant(), false)
	assertEquals(t, Term{IntCoeff(-1), "x"}.IsConstant(), false)
	assertEquals(t, Term{IntCoeff(2), "x"}.IsConstant(), false)
	assertEquals(t, Term{IntCoeff(0), ""}.IsConstant(), true)
	assertEquals(t, Term{IntCoeff(1), ""}.IsConstant(), true)
	assertEquals(t, Term{IntCoeff(-1), ""}.IsConstant(), true)
	assertEquals(t, Term{IntCoeff(2), ""}.IsConstant(), true)
}

func TestFormat(t *testing.T) {
	assertEquals(t, Term{IntCoeff(0), "x"}.Format(true), "")
	assertEquals(t, Term{IntCoeff(1), "x"}.Format(true), "x")
	assertEquals(t, Term{IntCoeff(-1), "x"}.Format(true), "- x")
	assertEquals(t, Term{IntCoeff(2), "x"}.Format(true), "2x")
	assertEquals(t, Term{IntCoeff(0), ""}.Format(true), "")
	assertEquals(t, Term{IntCoeff(1), ""}.Format(true), "1")
	assertEquals(t, Term{IntCoeff(-1), ""}.Format(true), "- 1")
	assertEquals(t, Term{IntCoeff(2), ""}.Format(true), "2")
	assertEquals(t, Term{IntCoeff(0), "x"}.Format(false), "")
	assertEquals(t, Term{IntCoeff(1), "x"}.Format(false), "+ x")
	assertEquals(t, Term{IntCoeff(-1), "x"}.Format(false), "- x")
	assertEquals(t, Term{IntCoeff(2), "x"}.Format(false), "+ 2x")
	assertEquals(t, Term{IntCoeff(0), ""}.Format(false), "")
	assertEquals(t, Term{IntCoeff(1), ""}.Format(false), "+ 1")
	assertEquals(t, Term{IntCoeff(-1), ""}.Format(false), "- 1")
	assertEquals(t, Term{IntCoeff(2), ""}.Format(false), "+ 2")
}
