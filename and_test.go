package gomda

import "testing"

func TestAnd_ShouldPassWithBoolArgs(t *testing.T) {
	if !And(true, true) {
		t.Error("true and true should return true")
	}

	if And(true, false) {
		t.Error("true and false should return false")
	}

	if And(false, false) {
		t.Error("false and false should return false")
	}

	if And(false, true) {
		t.Error("false and true should return false")
	}
}

func TestAnd_ShouldReturnFalseWithNilArgs(t *testing.T) {
	if And(true, nil) {
		t.Error("true and nil should return false")
	}

	if And(nil, true) {
		t.Error("nil and true should return false")
	}
}

func TestAnd_ShouldReturnFalseWithZeroValues(t *testing.T) {
	if And(true, 0) {
		t.Error("true and int should return false")
	}

	if And(true, 0.0) {
		t.Error("true and float should return false")
	}

	if And(true, "") {
		t.Error("true and string should return false")
	}

	if And(true, []int{}) {
		t.Error("true and slice should return false")
	}

	if And(true, [0]int{}) {
		t.Error("true and array should return false")
	}

	if And(true, map[int]string{}) {
		t.Error("true and map should return false")
	}

	type a struct {
		FieldOne string
		FieldTwo int
	}
	if And(true, a{}) {
		t.Error("true and struct should return false")
	}
}

func TestAnd_ShouldReturnTrueWithNonZeroValues(t *testing.T) {
	if !And(true, 1) {
		t.Error("true and int should return true")
	}

	if !And(true, 0.1) {
		t.Error("true and float should return true")
	}

	if !And(true, "hello") {
		t.Error("true and string should return true")
	}

	if !And(true, []int{1}) {
		t.Error("true and slice should return true")
	}

	if !And(true, [1]int{1}) {
		t.Error("true and array should return true")
	}

	if !And(true, map[int]string{1: "hello"}) {
		t.Error("true and map should return true")
	}

	type a struct {
		FieldOne string
		FieldTwo int
	}
	if !And(true, a{"hello", 1}) {
		t.Error("true and struct should return true")
	}
}
