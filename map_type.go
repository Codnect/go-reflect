package reflect

import "reflect"

type MapType struct {
	typ reflect.Type
}

func newMapType(p reflect.Type) MapType {
	return MapType{
		typ: p,
	}
}

func (t MapType) GetKeyType() Type {
	return Type{}
}

func (t MapType) GetValueType() Type {
	return Type{}
}
