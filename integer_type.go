package reflect

import "reflect"

type IntegerType struct {
	typ reflect.Type
}

func newIntegerType(p reflect.Type) IntegerType {
	return IntegerType{
		typ: p,
	}
}

func (t IntegerType) IsSignedInteger() bool {
	kind := t.typ.Kind()
	return kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Int
}

func (t IntegerType) IsUnsignedInteger() bool {
	kind := t.typ.Kind()
	return kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Uint
}
