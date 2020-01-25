package reflect

import "reflect"

type ArrayType struct {
	typ reflect.Type
}

func newArrayType(p reflect.Type) ArrayType {
	return ArrayType{
		typ: p,
	}
}

func (t ArrayType) GetElementType() Type {
	return newTypeWithType(t.typ.Elem())
}
