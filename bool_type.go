package reflect

import "reflect"

type BoolType struct {
	typ reflect.Type
}

func newBoolType(p reflect.Type) BoolType {
	return BoolType{
		typ: p,
	}
}
