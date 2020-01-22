package reflect

import (
	"reflect"
)

type Type struct {
	typ reflect.Type
}

func newType(p reflect.Type) Type {
	return Type{
		typ: p,
	}
}

func (t Type) GetName() string {
	return t.typ.String()
}

func (t Type) GetSimpleName() string {
	return t.typ.Name()
}

func (t Type) GetPackagePath() string {
	return t.typ.PkgPath()
}

func (t Type) IsTag() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		tagType := reflect.TypeOf((*Tag)(nil)).Elem()
		return typ.Implements(tagType)
	}
	return false
}

func (t Type) IsPointer() bool {
	return t.typ.Kind() == reflect.Ptr
}

func (t Type) IsInterface() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Interface
}

func (t Type) IsSlice() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Slice
}

func (t Type) IsStruct() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Struct
}

func (t Type) IsFunction() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Func
}

func (t Type) IsAssignableTo(p Type) bool {
	return false
}

func (t Type) Struct() (Struct, bool) {
	if !t.IsStruct() {
		return Struct{}, false
	}
	return newStruct(t.typ), true
}

func (t Type) Interface() (Interface, bool) {
	if !t.IsInterface() {
		return Interface{}, false
	}
	return newInterface(t.typ), true
}

func (t Type) Function() (Function, bool) {
	if !t.IsFunction() {
		return Function{}, false
	}
	return Function{}, true
}
