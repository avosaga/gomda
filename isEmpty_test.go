package gomda

import "testing"

func TestIsEmpty_ZeroValues(t *testing.T) {
	if !IsEmpty(0) {
		t.Error("Expected to return true if int is zero value")
	}

	if !IsEmpty(0.0) {
		t.Error("Expected to return true if float is zero value")
	}

	if !IsEmpty("") {
		t.Error("Expected to return true if string is zero value")
	}

	if !IsEmpty(false) {
		t.Error("Expected to return true if bool is zero value")
	}

	if !IsEmpty([]int{}) {
		t.Error("Expected to return true if slice is zero value")
	}

	if !IsEmpty([0]int{}) {
		t.Error("Expected to return true if array is zero value")
	}

	if !IsEmpty(make(map[int]int)) {
		t.Error("Expected to return true if map is zero value")
	}

	type A struct {
		FieldOne string
		FieldTwo int
	}
	if !IsEmpty(A{}) {
		t.Error("Expected to return true if struct is zero value")
	}

	var x int
	if !IsEmpty(&x) {
		t.Error("Expected to return true if pointer is zero value")
	}
}

func TestIsEmpty_NoZeroValues(t *testing.T) {
	if IsEmpty(1) {
		t.Error("Expected to return false if int is not zero value")
	}

	if IsEmpty(0.1) {
		t.Error("Expected to return false if float is not zero value")
	}

	if IsEmpty("hello") {
		t.Error("Expected to return false if string is not zero value")
	}

	if IsEmpty(true) {
		t.Error("Expected to return false if bool is not zero value")
	}

	if IsEmpty([]int{1}) {
		t.Error("Expected to return false if slice is not zero value")
	}

	if IsEmpty([1]int{1}) {
		t.Error("Expected to return false if array is not zero value")
	}

	m := make(map[int]int)
	m[1] = 1
	if IsEmpty(m) {
		t.Error("Expected to return false if map is not zero value")
	}

	type A struct {
		FieldOne string
		FieldTwo int
	}
	if IsEmpty(A{"hello", 1}) {
		t.Error("Expected to return false if struct is not zero value")
	}

	x := 1
	if IsEmpty(&x) {
		t.Error("Expected to return false if pointer is not zero value")
	}
}

func TestIsEmpty_ShouldReturnFalseIfValueIsNil(t *testing.T) {
	if !IsEmpty(nil) {
		t.Fail()
	}
}
