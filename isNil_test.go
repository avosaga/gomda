package gomda

import "testing"

func TestIsNil_ShouldReturnTrueIfValueIsNil(t *testing.T) {
	if !IsNil(nil) {
		t.Fail()
	}
}

func TestIsNil_ShouldReturnFalseIfValueIsNotNil(t *testing.T) {
	if IsNil(1) {
		t.Fail()
	}
}
