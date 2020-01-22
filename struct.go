package reflect

import (
	"errors"
	"fmt"
	"github.com/codnect/go-one"
	"reflect"
)

type Struct struct {
	typ reflect.Type
}

func newStruct(p reflect.Type) Struct {
	return Struct{
		typ: p,
	}
}

func (t Struct) GetDeclaredFieldByIndex(index int) (Field, error) {
	return t.privateGetFieldByIndex(index)
}

func (t Struct) GetFieldByName(name string) (Field, error) {
	return t.privateGetField(true, name)
}

func (t Struct) GetFields() []Field {
	return t.privateGetFields(true)
}

func (t Struct) GetDeclaredFieldByName(name string) (Field, error) {
	return t.privateGetField(false, name)
}

func (t Struct) GetDeclaredFields() []Field {
	return t.privateGetFields(false)
}

func (t Struct) privateGetField(exportedOnly bool, name string) (Field, error) {
	structField, result := getStructFieldByName(t.typ, name)
	if !result {
		return Field{}, errors.New(fmt.Sprintf("field named %s not found", name))
	}
	isExported := isExportedField(structField)
	if exportedOnly && !isExported {
		return Field{}, errors.New(fmt.Sprintf("field named %s is not exported", name))
	}
	return newField(structField), nil
}

func (t Struct) privateGetFieldByIndex(index int) (Field, error) {
	numField := getStructNumField(t.typ)
	if index >= 0 && index < numField {
		structField := getStructFieldByIndex(t.typ, index)
		field := newField(structField)
		return field, nil
	}
	return Field{}, errors.New(fmt.Sprint("field indexed at %i not found", index))
}

func (t Struct) privateGetFields(exportedOnly bool) []Field {
	fields := make([]Field, 0)
	for index := 0; index < getStructNumField(t.typ); index++ {
		structField := getStructFieldByIndex(t.typ, index)
		isExported := isExportedField(structField)
		if exportedOnly && !isExported {
			continue
		}
		field := newField(structField)
		fields = append(fields, field)
	}
	return fields
}

/*
func (t StructType) NewInstance() one.One {
	return reflect.New(getType(t.one)).Elem().Interface()
}

func (t StructType) NewInstancePointer() one.One {
	return reflect.New(getType(t.one)).Interface()
}*/

func (t Struct) GetMethodByName(name string) (Method, error) {
	return t.privateGetMethod(true, name)
}

func (t Struct) GetMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t Struct) GetMethods() []Method {
	return t.privateGetMethods(true)
}

func (t Struct) GetDeclaredMethodByName(name string) (Method, error) {
	return t.privateGetMethod(false, name)
}

func (t Struct) GetDeclaredMethods() []Method {
	return t.privateGetMethods(false)
}

func (t Struct) GetDeclaredMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t Struct) privateGetMethod(exportedOnly bool, name string) (Method, error) {
	structMethod, result := getMethodByName(t.typ, name)
	if !result {
		return Method{}, errors.New(fmt.Sprintf("method named %s not found", name))
	}
	isExported := isExportedMethod(structMethod)
	if exportedOnly && !isExported {
		return Method{}, errors.New(fmt.Sprintf("method named %s is not exported", name))
	}
	return newMethod(structMethod), nil
}

func (t Struct) privateGetMethods(exportedOnly bool) []Method {
	methods := make([]Method, 0)
	for index := 0; index < getNumMethod(t.typ); index++ {
		method := getMethodByIndex(t.typ, index)
		isExported := isExportedMethod(method)
		if exportedOnly && !isExported {
			continue
		}
		methods = append(methods, newMethod(method))
	}
	return methods
}
