package gomda

import "testing"

func TestAdjust_ShouldReturnErrorIfListIsNotAnArrayNorSlice(t *testing.T) {
	err1 := Adjust(0, nil, 1)
	err2 := Adjust(0, nil, nil)

	if err1 == nil || err2 == nil {
		t.Fail()
	}
}

func TestAdjust_ShouldReturnErrorIfApplyFunctionIsNotAFunction(t *testing.T) {
	err := Adjust(0, 1, []int{1})

	if err == nil {
		t.Fail()
	}
}

func TestAdjust_ShouldReturnErrorIfApplyFunctionDoesNotMatchListType(t *testing.T) {
	err1 := Adjust(0, func(i float32) int { return 2 }, []int{1})
	err2 := Adjust(0, func(i int) float32 { return 2.0 }, []int{1})

	if err1 == nil || err2 == nil {
		t.Fail()
	}
}

func TestAdjust_ShouldReturnErrorIfIndexIsOutOfBounds(t *testing.T) {
	fn := func(i int) int { return 2 }
	list := []int{1}
	err1 := Adjust(3, fn, list)
	err2 := Adjust(-3, fn, list)

	if err1 == nil || err2 == nil {
		t.Fail()
	}
}

func TestAdjust_ShouldNotReturnErrorAndAdjustValue(t *testing.T) {
	fn := func(i int) int { return 2 }
	index := 2
	list1 := []int{1, 2, 3, 4}
	list2 := []int{1, 2, 3, 4}
	err1 := Adjust(index, fn, list1)
	err2 := Adjust(-2, fn, list2)

	if err1 != nil || err2 != nil {
		t.Fail()
	}

	if list1[index] != 2 || list2[index] != 2 {
		t.Fail()
	}
}
