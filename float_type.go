package reflect

import "reflect"

type FloatType struct {
	typ reflect.Type
}

func newFloatType(p reflect.Type) FloatType {
	return FloatType{
		typ: p,
	}
}

func (t FloatType) IsFloat32() bool {
	return t.typ.Kind() == reflect.Float32
}

func (t FloatType) IsFloat64() bool {
	return t.typ.Kind() == reflect.Float64
}
