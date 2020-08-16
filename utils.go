package gomda

import "reflect"

type GomdaError struct {
	message string
}

func (e *GomdaError) Error() string {
	return e.message
}

func ReadValue(value reflect.Value) interface{} {
	switch value.Kind() {
	case Bool:
		return value.Bool()
	case Int:
		return value.Int()
	case Int8:
		return value.Int()
	case Int16:
		return value.Int()
	case Int32:
		return value.Int()
	case Int64:
		return value.Int()
	case Uint:
		return value.Uint()
	case Uint8:
		return value.Uint()
	case Uint16:
		return value.Uint()
	case Uint32:
		return value.Uint()
	case Uint64:
		return value.Uint()
	case Float32:
		return value.Float()
	case Float64:
		return value.Float()
	case Complex64:
		return value.Complex()
	case Complex128:
		return value.Complex()
	case String:
		return value.String()
	}

	return nil
}

func ValidateFunction(fn interface{}, expectedIn []reflect.Kind, expectedOut []reflect.Kind) error {
	if !Is(reflect.Func, fn) {
		return &GomdaError{"arg: 'fn' must be a valid function with matching types of arguments and return values!"}
	}

	fnType := reflect.TypeOf(fn)
	numIn := fnType.NumIn()
	numOut := fnType.NumOut()
	expIn := 0
	expOut := 0

	if !IsNil(expectedIn) {
		expIn = len(expectedIn)
	}

	if !IsNil(expectedOut) {
		expOut = len(expectedOut)
	}

	if numIn != expIn {
		return &GomdaError{"Function number of in-args do not match expected ones"}
	}

	if numOut != expOut {
		return &GomdaError{"Function number of out-args do not match expected ones"}
	}

	for i := 0; i < numIn; i++ {
		if fnType.In(i).Kind() != expectedIn[i] {
			return &GomdaError{"Function in-args types do not match expected types"}
		}
	}

	for i := 0; i < numOut; i++ {
		if fnType.Out(i).Kind() != expectedOut[i] {
			return &GomdaError{"Function out-args types do not match expected types"}
		}
	}

	return nil
}

const (
	String     = reflect.String
	Float64    = reflect.Float64
	Float32    = reflect.Float32
	Int        = reflect.Int
	Int8       = reflect.Int8
	Int16      = reflect.Int16
	Int32      = reflect.Int32
	Int64      = reflect.Int64
	Uint       = reflect.Uint
	Uint8      = reflect.Uint8
	Uint16     = reflect.Uint16
	Uint32     = reflect.Uint32
	Uint64     = reflect.Uint64
	Bool       = reflect.Bool
	Complex64  = reflect.Complex64
	Complex128 = reflect.Complex128
	Struct     = reflect.Struct
	Array      = reflect.Array
	Slice      = reflect.Slice
	Map        = reflect.Map
	Pointer    = reflect.Ptr
)

func IsArrayOrSlice(v interface{}) bool {
	return Is(Array, v) || Is(Slice, v)
}

func IsNilOrEmpty(v interface{}) bool {
	return IsNil(v) || IsEmpty(v)
}
