package gomda

import "reflect"

func And(a, b interface{}) bool {
	return test(a) && test(b)
}

func test(v interface{}) bool {
	if Is(Bool, v) {
		return reflect.ValueOf(v).Bool()
	}

	return !IsEmpty(v)
}
