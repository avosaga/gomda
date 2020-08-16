package gomda

import "reflect"

func IsEmpty(value interface{}) bool {
	if IsNil(value) {
		return true
	}

	val := reflect.ValueOf(value)

	switch reflect.TypeOf(value).Kind() {
	case Slice:
		return val.Len() == 0
	case Map:
		return len(val.MapKeys()) == 0
	case Pointer:
		return val.Elem().IsZero()
	}

	return val.IsZero()
}
