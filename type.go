package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type Type struct {
	one one.One
}

func newType(one one.One) Type {
	return Type{
		one: one,
	}
}

func (t Type) GetName() string {
	return getType(t.one).String()
}

func (t Type) GetSimpleName() string {
	return getType(t.one).Name()
}

func (t Type) GetPackagePath() string {
	return getType(t.one).PkgPath()
}

func (t Type) IsTag() bool {
	if _, ok := t.one.(Tag); ok {
		return true
	}
	return false
}

func (t Type) IsPointer() bool {
	return reflect.ValueOf(t.one).Kind() == reflect.Ptr
}

func (t Type) IsInterface() bool {
	return getType(t.one).Kind() == reflect.Interface
}

func (t Type) IsSlice() bool {
	return getType(t.one).Kind() == reflect.Slice
}

func (t Type) IsStruct() bool {
	return getValue(t.one).Kind() == reflect.Struct
}

func (t Type) IsAssignableTo(p Type) bool {
	return false
}

func (t Type) StructType() (StructType, bool) {
	if !t.IsStruct() {
		return StructType{}, false
	}
	return newStructType(t.one), true
}
