package gomda

import "reflect"

func Update(index int, value interface{}, list interface{}) error {
	if IsNil(value) {
		return &GomdaError{"Value should not be nil"}
	}

	fn := reflect.FuncOf(
		[]reflect.Type{reflect.TypeOf(value)},
		[]reflect.Type{reflect.TypeOf(value)},
		false,
	)

	function := func(args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(value)}
	}

	return Adjust(index, reflect.MakeFunc(fn, function).Interface(), list)
}
