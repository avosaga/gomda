package gomda

import "testing"

func TestAllPass_ShouldPass(t *testing.T) {
	isGreaterThanZero := func(n int) bool {
		return n > 0
	}

	isLessThanFive := func(n int) bool {
		return n < 5
	}

	test := AllPass([]interface{}{isGreaterThanZero, isLessThanFive})

	if !test(2) {
		t.Fail()
	}

	if test(6) || test(-1) {
		t.Fail()
	}
}

//TODO: missing tests
