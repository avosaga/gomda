package gomda

import "reflect"

func Equals(a interface{}, b interface{}) bool {
	if IsNil(a) && IsNil(b) {
		return true
	}

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	var aValue, bValue = reflect.ValueOf(a), reflect.ValueOf(b)

	switch reflect.TypeOf(a).Kind() {
	case Pointer:
		return equalsPointer(aValue, bValue)
	case Struct:
		return equalsStruct(aValue, bValue)
	case Array:
		return equalsCollection(aValue, bValue)
	case Slice:
		return equalsCollection(aValue, bValue)
	case Map:
		return equalsMap(aValue, bValue)
	default:
		return ReadValue(aValue) == ReadValue(bValue)
	}
}

func EqualsTo(value interface{}) func(interface{}) bool {
	return func(other interface{}) bool {
		return Equals(value, other)
	}
}

func equalsStruct(s1 reflect.Value, s2 reflect.Value) bool {
	if s1.NumField() != s2.NumField() {
		return false
	}

	for i := 0; i < s1.NumField(); i++ {
		name1 := s1.Type().Field(i).Name
		val1 := s1.Field(i).Interface()
		fieldFound := false

		for j := 0; j < s2.NumField(); j++ {
			name2 := s2.Type().Field(i).Name

			if name2 == name1 {
				fieldFound = true

				if !Equals(val1, s2.Field(i).Interface()) {
					return false
				}
			}
		}

		if !fieldFound {
			return false
		}
	}

	return true
}

func equalsCollection(a1 reflect.Value, a2 reflect.Value) bool {
	if a1.Len() != a2.Len() {
		return false
	}

	for i := 0; i < a1.Len(); i++ {
		aValue := a1.Index(i)
		bValue := a2.Index(i)

		if !Equals(aValue.Interface(), bValue.Interface()) {
			return false
		}
	}

	return true
}

func equalsMap(a1 reflect.Value, a2 reflect.Value) bool {
	a1Keys := a1.MapKeys()
	a2Keys := a2.MapKeys()

	if len(a1Keys) != len(a2Keys) {
		return false
	}

	for i := 0; i < len(a1Keys); i++ {
		currentKey := a1Keys[i]
		a1KeyValue := a1.MapIndex(currentKey)
		a2KeyValue := a2.MapIndex(currentKey)

		if !a2KeyValue.IsValid() || !Equals(a1KeyValue.Interface(), a2KeyValue.Interface()) {
			return false
		}
	}

	return true
}

func equalsPointer(a1 reflect.Value, a2 reflect.Value) bool {
	return Equals(a1.Elem().Interface(), a2.Elem().Interface())
}
