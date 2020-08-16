package gomda

import (
	"testing"
)

func assertError(err error, t *testing.T) {
	if err == nil {
		t.Error(err)
	}
}

func TestForEach_ShouldIterateListWithValidArgs(t *testing.T) {
	list := []int{1, 2, 3}
	counter := 0
	err := ForEach(list, func(v int, i int) bool {
		counter++
		return false
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != len(list) {
		t.Error("should have iterated the complete list")
	}
}

func TestForEach_ShouldReturnErrorIfCallbackDefinitionIsInvalidForList(t *testing.T) {
	list := []int{1, 2, 3}
	incorrectValueTypeArgErr := ForEach(list, func(value float32, index int) bool {
		return false
	})
	incorrectIndexTypeArgErr := ForEach(list, func(value int, index float32) bool {
		return false
	})
	incorrectReturnedValueTypeErr := ForEach(list, func(value int, index int) string {
		return ""
	})
	missingReturnedValueErr := ForEach(list, func(value int, index int) {})
	nilFirstArgErr := ForEach(nil, func(value int, index int) bool { return false })
	nilSecondArgErr := ForEach(list, nil)

	assertError(incorrectValueTypeArgErr, t)
	assertError(incorrectIndexTypeArgErr, t)
	assertError(incorrectReturnedValueTypeErr, t)
	assertError(missingReturnedValueErr, t)
	assertError(nilFirstArgErr, t)
	assertError(nilSecondArgErr, t)
}

func TestForEach_ShouldBreakIfReturnTrue(t *testing.T) {
	list := []int{1, 2, 3}
	counter := 0
	err := ForEach(list, func(v int, i int) bool {
		counter++
		return v == 1
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != 1 {
		t.Error("should have broken loop")
	}
}

func TestForEach_ShouldNotIterateIfListIsEmpty(t *testing.T) {
	counter := 0

	err := ForEach([]int{}, func(v int, i int) bool {
		counter++
		return false
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != 0 {
		t.Error("should not loop")
	}
}

func TestForEach_ShouldIterateMapWithValidArgs(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	counter := 0

	err := ForEach(m, func(key int, value string) bool {
		counter++
		switch key {
		case 1:
			return value != "one"
		case 2:
			return value != "two"
		case 3:
			return value != "three"
		}
		return false
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != len(m) {
		t.Error("should have iterated the complete map")
	}
}

func TestForEach_ShouldReturnErrorIfCallbackDefinitionIsInvalidForMap(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	incorrectValueTypeArgErr := ForEach(m, func(key int, value float32) bool {
		return false
	})
	incorrectKeyTypeArgErr := ForEach(m, func(key string, value string) bool {
		return false
	})
	incorrectReturnedValueTypeErr := ForEach(m, func(key int, value string) string {
		return ""
	})
	missingReturnedValueErr := ForEach(m, func(key int, value string) {})
	nilFirstArgErr := ForEach(nil, func(key int, value string) bool { return false })
	nilSecondArgErr := ForEach(m, nil)

	assertError(incorrectValueTypeArgErr, t)
	assertError(incorrectKeyTypeArgErr, t)
	assertError(incorrectReturnedValueTypeErr, t)
	assertError(missingReturnedValueErr, t)
	assertError(nilFirstArgErr, t)
	assertError(nilSecondArgErr, t)
}

func TestForEach_ShouldNotIterateIfMapIsEmpty(t *testing.T) {
	counter := 0

	err := ForEach(map[int]string{}, func(k int, v string) bool {
		counter++
		return false
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != 0 {
		t.Error("should not loop")
	}
}

func TestForEach_ShouldIterateIntWithValidArgs(t *testing.T) {
	counter := 0
	err := ForEach(10, func(val int) bool {
		counter++
		return false
	})

	if err != nil {
		t.Error("should not return error")
	}

	if counter != 10 {
		t.Error("should not loop")
	}
}
