package reflect

import "reflect"

var intTypes = map[reflect.Kind]bool{
	reflect.Int:    true,
	reflect.Int8:   true,
	reflect.Int16:  true,
	reflect.Int32:  true,
	reflect.Int64:  true,
	reflect.Uint:   true,
	reflect.Uint8:  true,
	reflect.Uint16: true,
	reflect.Uint32: true,
	reflect.Uint64: true,
}

type NumberType struct {
	typ reflect.Type
}

func newNumberType(p reflect.Type) NumberType {
	return NumberType{
		typ: p,
	}
}

func (t NumberType) IsInteger() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if _, ok := intTypes[typ.Kind()]; ok {
		return true
	}
	return false
}

func (t NumberType) IsFloat() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	kind := typ.Kind()
	return kind == reflect.Float32 || kind == reflect.Float64
}

func (t NumberType) IsComplex() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	kind := typ.Kind()
	return kind == reflect.Complex64 || kind == reflect.Complex128
}
