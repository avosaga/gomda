package gomda

import (
	"reflect"
)

func AllPass(predicates []interface{}) func(interface{}) bool {
	return func(value interface{}) bool {
		result := true
		err := ForEach(predicates, func(predicate interface{}, i int) bool {
			val := reflect.ValueOf(value)
			err := ValidateFunction(
				predicate,
				[]reflect.Kind{val.Kind()},
				[]reflect.Kind{Bool},
			)

			if !IsNil(err) {
				result = false
				return true
			}

			results := reflect.ValueOf(predicate).Call([]reflect.Value{val})
			result = results[0].Bool() && result

			return false
		})

		return IsNil(err) && result
	}
}
