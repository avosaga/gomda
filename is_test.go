package gomda

import (
	"math/rand"
	"testing"
)

func TestIs_ShouldReturnTrueIfValueMatchesType(t *testing.T) {
	if !Is(Int, rand.Int()) {
		t.Fail()
	}
}

func TestIs_ShouldReturnFalseIfValueDoesNotMatchType(t *testing.T) {
	if Is(Int, "1") {
		t.Fail()
	}
}

func TestIs_ShouldReturnFalseIfValueIsNil(t *testing.T) {
	if Is(String, nil) {
		t.Fail()
	}
}
