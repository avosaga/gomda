package gomda

import "testing"

func TestAll_ShouldReturnFalseIfValueIsNotAnArrayNorSlice(t *testing.T) {
	if All(func(i int) bool { return false })(1) {
		t.Fail()
	}
}

func TestAll_ShouldReturnFalseIfValueIsNil(t *testing.T) {
	if All(func(i int) bool { return false })(nil) {
		t.Fail()
	}
}

func TestAll_ShouldReturnFalseIfValueIsEmpty(t *testing.T) {
	if All(func(i int) bool { return false })([]int{}) {
		t.Fail()
	}
}

func TestAll_ShouldReturnFalseIfPredicateIsNotAFunction(t *testing.T) {
	if All(1)([]int{1}) {
		t.Fail()
	}
}

func TestAll_ShouldReturnFalseIfPredicateFunctionArgDoesNotMatchValueType(t *testing.T) {
	if All(func(i float64) bool { return false })([]int{1}) {
		t.Fail()
	}
}

func TestAll_ShouldReturnTrueIfAllElementsOfTheListMatchThePredicate(t *testing.T) {
	predicate := func(n int) bool {
		return n == 1
	}

	if !All(predicate)([]int{1, 1, 1, 1, 1, 1}) {
		t.Fail()
	}
}

func TestAll_ShouldReturnTrueIfAllElementsOfTheListMatchThePredicateUsingStruct(t *testing.T) {
	type A struct {
		FieldOne int
		FieldTwo string
	}
	predicate := func(a A) bool {
		return a.FieldOne == 1 && a.FieldTwo == "hello"
	}

	if !All(predicate)([]A{{1, "hello"}, {1, "hello"}}) {
		t.Fail()
	}
}

func TestAll_ShouldReturnFalseIfAtLeastOneElementOfTheListDoesNotMatchThePredicate(t *testing.T) {
	predicate := func(n int) bool {
		return n == 1
	}

	if All(predicate)([]int{1, 1, 1, 1, 2, 1}) {
		t.Fail()
	}
}
