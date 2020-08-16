package gomda

import "reflect"

func Is(kind reflect.Kind, value interface{}) bool {
	return !IsNil(value) && reflect.TypeOf(value).Kind() == kind
}
