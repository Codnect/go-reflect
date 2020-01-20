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
	return getValue(t.one).Kind() == reflect.Interface
}

func (t Type) IsStruct() bool {
	return getValue(t.one).Kind() == reflect.Struct
}

func (t Type) IsAssignableTo(p Type) bool {
	return false
}

func (t Type) GetDeclaredFieldByIndex(index int) (Field, bool) {
	return t.privateGetFieldByIndex(index)
}

func (t Type) GetFieldByName(name string) (Field, bool) {
	return t.privateGetField(true, name)
}

func (t Type) GetFields() []Field {
	return t.privateGetFields(true)
}

func (t Type) GetDeclaredFieldByName(name string) (Field, bool) {
	return t.privateGetField(false, name)
}

func (t Type) GetDeclaredFields() []Field {
	return t.privateGetFields(false)
}

func (t Type) privateGetField(exportedOnly bool, name string) (Field, bool) {
	structField, result := getStructFieldByName(t.one, name)
	if !result {
		return Field{}, false
	}
	isExported := isExportedField(structField)
	if exportedOnly && !isExported {
		return Field{}, false
	}
	field := newField(structField)
	return field, true
}

func (t Type) privateGetFieldByIndex(index int) (Field, bool) {
	numField := getNumField(t.one)
	if index >= 0 && index < numField {
		structField := getStructFieldByIndex(t.one, index)
		field := newField(structField)
		return field, true
	}
	return Field{}, false
}

func (t Type) privateGetFields(exportedOnly bool) []Field {
	fields := make([]Field, 0)
	for index := 0; index < getNumField(t.one); index++ {
		structField := getStructFieldByIndex(t.one, index)
		isExported := isExportedField(structField)
		if exportedOnly && !isExported {
			continue
		}
		field := newField(structField)
		fields = append(fields, field)
	}
	return fields
}
