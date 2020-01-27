package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type MapType struct {
	typ reflect.Type
}

func newMapType(p reflect.Type) MapType {
	return MapType{
		typ: p,
	}
}

func (t MapType) GetKeyType() Type {
	return newTypeWithType(t.typ.Key())
}

func (t MapType) GetValueType() Type {
	return newTypeWithType(t.typ.Elem())
}

func (t MapType) NewInstance() one.One {
	return reflect.MakeMap(reflect.MapOf(t.typ.Key(), t.typ.Elem()))
}
