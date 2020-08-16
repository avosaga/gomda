package gomda

import "reflect"

func All(fn interface{}) func(interface{}) bool {
	return func(list interface{}) bool {
		if !IsArrayOrSlice(list) || IsEmpty(list) {
			return false
		}

		l := reflect.ValueOf(list)
		err := ValidateFunction(fn, []reflect.Kind{l.Index(0).Kind()}, []reflect.Kind{Bool})

		if !IsNil(err) {
			return false
		}

		function := reflect.ValueOf(fn)

		for i := 0; i < l.Len(); i++ {
			valueAtIndex := l.Index(i)
			result := function.Call([]reflect.Value{valueAtIndex})

			if !result[0].Bool() {
				return false
			}
		}

		return true
	}
}
