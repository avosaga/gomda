package gomda

import "reflect"

func ForEach(a interface{}, fn interface{}) error {
	if IsNil(a) {
		return &GomdaError{"ForEach is not expecting nil args!"}
	}

	if IsEmpty(a) {
		return nil
	}

	switch reflect.TypeOf(a).Kind() {
	case Array:
		return iterateList(a, fn)
	case Slice:
		return iterateList(a, fn)
	case Map:
		return iterateMap(a, fn)
	case Int:
		return iterateInt(a, fn)
	}

	return &GomdaError{"Only Array ,Slice, Map and Int types can are accepted!"}
}

func iterateList(list interface{}, fn interface{}) error {
	l := reflect.ValueOf(list)

	if err := ValidateFunction(fn, []reflect.Kind{l.Index(0).Kind(), Int}, []reflect.Kind{Bool}); err != nil {
		return err
	}

	function := reflect.ValueOf(fn)

	for i := 0; i < l.Len(); i++ {
		r := function.Call([]reflect.Value{l.Index(i), reflect.ValueOf(i)})

		if r[0].Bool() {
			break
		}
	}

	return nil
}

func iterateMap(mapValue interface{}, fn interface{}) error {
	m := reflect.ValueOf(mapValue)
	keys := m.MapKeys()

	if err := ValidateFunction(fn, []reflect.Kind{keys[0].Kind(), m.MapIndex(keys[0]).Kind()}, []reflect.Kind{Bool}); err != nil {
		return err
	}

	function := reflect.ValueOf(fn)

	for i := 0; i < len(keys); i++ {
		r := function.Call([]reflect.Value{keys[i], m.MapIndex(keys[i])})

		if r[0].Bool() {
			break
		}
	}

	return nil
}

func iterateInt(val interface{}, fn interface{}) error {
	n := int(reflect.ValueOf(val).Int())

	if err := ValidateFunction(fn, []reflect.Kind{Int}, []reflect.Kind{Bool}); err != nil {
		return err
	}

	function := reflect.ValueOf(fn)

	for i := 0; i < n; i++ {
		r := function.Call([]reflect.Value{reflect.ValueOf(i)})

		if r[0].Bool() {
			break
		}
	}

	return nil
}
