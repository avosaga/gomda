package gomda

import (
	"reflect"
	"testing"
)

func TestValidateFunction_ShouldNotReturnError(t *testing.T) {
	fn := func(a int, b string) bool { return false }
	err := ValidateFunction(fn, []reflect.Kind{reflect.Int, reflect.String}, []reflect.Kind{reflect.Bool})

	if err != nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldNotReturnErrorWhenPassingNilForExpectedInArgsAndOutArgs(t *testing.T) {
	err := ValidateFunction(func() {}, nil, nil)

	if err != nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldReturnErrorWhenPassingADifferentType(t *testing.T) {
	err := ValidateFunction(1, nil, nil)

	if err == nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldReturnErrorWhenExpectedNumberOfInArgsAreDifferent(t *testing.T) {
	fn := func(a int) {}
	err := ValidateFunction(fn, []reflect.Kind{reflect.Int, reflect.String}, nil)

	if err == nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldReturnErrorWhenExpectedNumberOfOutArgsAreDifferent(t *testing.T) {
	fn := func() bool { return false }
	err := ValidateFunction(fn, nil, []reflect.Kind{reflect.Int, reflect.String})

	if err == nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldReturnErrorWhenExpectedInArgsTypesAreDifferent(t *testing.T) {
	fn := func(a int) {}
	err := ValidateFunction(fn, []reflect.Kind{reflect.String}, nil)

	if err == nil {
		t.Fail()
	}
}

func TestValidateFunction_ShouldReturnErrorWhenExpectedOutArgsTypesAreDifferent(t *testing.T) {
	fn := func() bool { return false }
	err := ValidateFunction(fn, nil, []reflect.Kind{reflect.String})

	if err == nil {
		t.Fail()
	}
}

func TestIsArrayOrSlice_ShouldReturnTrueIfValueIsAnArrayOrSlice(t *testing.T) {
	if !IsArrayOrSlice([]int{1}) || !IsArrayOrSlice([1]int{1}) {
		t.Fail()
	}
}

func TestIsArrayOrSlice_ShouldReturnFalseIfValueIsNotAnArrayNorSlice(t *testing.T) {
	if IsArrayOrSlice("hello") {
		t.Fail()
	}
}

func TestIsNilOrEmpty_ShouldReturnTrueIfValueIsNilOrEmpty(t *testing.T) {
	if !IsNilOrEmpty(nil) || !IsNilOrEmpty(0) {
		t.Fail()
	}
}

func TestIsNilOrEmpty_ShouldReturnFalseIfValueIsNotNilNorEmpty(t *testing.T) {
	if IsNilOrEmpty("hello") || IsNilOrEmpty(1) {
		t.Fail()
	}
}
