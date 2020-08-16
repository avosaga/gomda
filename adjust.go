package gomda

import "reflect"

func Adjust(index int, fn interface{}, list interface{}) error {
	if !IsArrayOrSlice(list) {
		return &GomdaError{"List must be an array or slice!"}
	}

	a := reflect.ValueOf(list)

	if index > a.Len()-1 || (index < 0 && index < a.Len()*-1) {
		return &GomdaError{"Index is out of bounds"}
	}

	i := index
	if index < 0 {
		i = a.Len() + index
	}

	valueAtIndex := a.Index(i)
	err := ValidateFunction(fn, []reflect.Kind{valueAtIndex.Kind()}, []reflect.Kind{valueAtIndex.Kind()})

	if !IsNil(err) {
		return err
	}

	function := reflect.ValueOf(fn)
	returnValues := function.Call([]reflect.Value{valueAtIndex})
	valueAtIndex.Set(returnValues[0])
	return nil
}
