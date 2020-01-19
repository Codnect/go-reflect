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
	return getValue(t.one).String()
}

func (t Type) GetSimpleName() string {
	return getValue(t.one).String()
}

func (t Type) IsPointer() bool {
	return isPointer(reflect.ValueOf(t.one))
}

func (t Type) IsInterface() bool {
	return isInterface(getValue(t.one))
}

func (t Type) IsStruct() bool {
	return isStruct(getValue(t.one))
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
	field := newField(structField.Name, structField.Anonymous, isExported)
	return field, true
}

func (t Type) privateGetFieldByIndex(index int) (Field, bool) {
	numField := getNumField(t.one)
	if index >= 0 && index < numField {
		structField := getStructFieldByIndex(t.one, index)
		isExported := isExportedField(structField)
		field := newField(structField.Name, structField.Anonymous, isExported)
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
		field := newField(structField.Name, structField.Anonymous, isExported)
		fields = append(fields, field)
	}
	return fields
}
