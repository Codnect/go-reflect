package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type SliceType struct {
	typ reflect.Type
}

func newSliceType(p reflect.Type) SliceType {
	return SliceType{
		typ: p,
	}
}

func (t SliceType) GetElementType() Type {
	return newTypeWithType(t.typ.Elem())
}

func (t SliceType) NewInstance() one.One {
	return nil
}
