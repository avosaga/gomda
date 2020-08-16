package gomda

import "testing"

func TestUpdate_ShouldNotReturnErrorAndAdjustValue(t *testing.T) {
	list := []int{1, 3}
	err := Update(0, 2, list)

	if err != nil {
		t.Error(err)
	}

	if list[0] != 2 {
		t.Error("Value at index 0 should have been updated")
	}
}

func TestUpdate_ShouldReturnErrorIfValueIsNil(t *testing.T) {
	err := Update(0, nil, []int{1})

	if err == nil {
		t.Fail()
	}
}
