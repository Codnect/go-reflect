package reflect

import (
	"reflect"
)

var numberTypes = map[reflect.Kind]bool{
	reflect.Int:        true,
	reflect.Int8:       true,
	reflect.Int16:      true,
	reflect.Int32:      true,
	reflect.Int64:      true,
	reflect.Uint:       true,
	reflect.Uint8:      true,
	reflect.Uint16:     true,
	reflect.Uint32:     true,
	reflect.Uint64:     true,
	reflect.Float32:    true,
	reflect.Float64:    true,
	reflect.Complex64:  true,
	reflect.Complex128: true,
}

type Type struct {
	typ reflect.Type
	val reflect.Value
}

func newType(typ reflect.Type, value reflect.Value) Type {
	return Type{
		typ: typ,
		val: value,
	}
}

func newTypeWithType(typ reflect.Type) Type {
	return Type{
		typ: typ,
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

func (t Type) IsArray() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Array
}

func (t Type) IsNumber() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if _, ok := numberTypes[typ.Kind()]; ok {
		return true
	}
	return false
}

func (t Type) IsBool() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Bool
}

func (t Type) IsMap() bool {
	typ := t.typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Map
}

func (t Type) IsAssignableTo(p Type) bool {
	return false
}

func (t Type) StructType() (StructType, bool) {
	if !t.IsStruct() {
		return StructType{}, false
	}
	return newStructType(t.typ), true
}

func (t Type) InterfaceType() (InterfaceType, bool) {
	if !t.IsInterface() {
		return InterfaceType{}, false
	}
	return newInterfaceType(t.typ), true
}

func (t Type) FunctionType() (FunctionType, bool) {
	if !t.IsFunction() {
		return FunctionType{}, false
	}
	return FunctionType{}, true
}

func (t Type) MapType() (MapType, bool) {
	if !t.IsMap() {
		return MapType{}, false
	}
	return newMapType(t.typ), true
}

func (t Type) ArrayType() (ArrayType, bool) {
	if !t.IsArray() {
		return ArrayType{}, false
	}
	return newArrayType(t.typ), true
}

func (t Type) NumberType() (NumberType, bool) {
	if !t.IsArray() {
		return NumberType{}, false
	}
	return newNumberType(t.typ), true
}

func (t Type) BoolType() (BoolType, bool) {
	if !t.IsArray() {
		return BoolType{}, false
	}
	return newBoolType(t.typ), true
}
