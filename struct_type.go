package reflect

import (
	"errors"
	"fmt"
	"github.com/codnect/go-one"
	"reflect"
)

type StructType struct {
	one one.One
}

func newStructType(one one.One) StructType {
	return StructType{
		one: one,
	}
}

func (t StructType) GetDeclaredFieldByIndex(index int) (Field, error) {
	return t.privateGetFieldByIndex(index)
}

func (t StructType) GetFieldByName(name string) (Field, error) {
	return t.privateGetField(true, name)
}

func (t StructType) GetFields() ([]Field, error) {
	return t.privateGetFields(true)
}

func (t StructType) GetDeclaredFieldByName(name string) (Field, error) {
	return t.privateGetField(false, name)
}

func (t StructType) GetDeclaredFields() ([]Field, error) {
	return t.privateGetFields(false)
}

func (t StructType) privateGetField(exportedOnly bool, name string) (Field, error) {
	structField, result := getStructFieldByName(t.one, name)
	if !result {
		return Field{}, errors.New(fmt.Sprintf("field named %s not found", name))
	}
	isExported := isExportedField(structField)
	if exportedOnly && !isExported {
		return Field{}, errors.New(fmt.Sprintf("field named %s is not exported", name))
	}
	return newField(structField), nil
}

func (t StructType) privateGetFieldByIndex(index int) (Field, error) {
	numField := getNumField(t.one)
	if index >= 0 && index < numField {
		structField := getStructFieldByIndex(t.one, index)
		field := newField(structField)
		return field, nil
	}
	return Field{}, errors.New(fmt.Sprint("field indexed at %i not found", index))
}

func (t StructType) privateGetFields(exportedOnly bool) ([]Field, error) {
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
	return fields, nil
}

func (t StructType) NewInstance() one.One {
	return reflect.New(getType(t.one)).Elem().Interface()
}

func (t StructType) NewInstancePointer() one.One {
	return reflect.New(getType(t.one)).Interface()
}

func (t StructType) GetMethodByName(name string) (Method, error) {
	return t.privateGetMethod(true, name)
}

func (t StructType) GetMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t StructType) GetMethods() ([]Method, error) {
	return t.privateGetMethods(true)
}

func (t StructType) GetDeclaredMethodByName(name string) (Method, error) {
	return t.privateGetMethod(false, name)
}

func (t StructType) GetDeclaredMethods() ([]Method, error) {
	return t.privateGetMethods(false)
}

func (t StructType) GetDeclaredMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t StructType) privateGetMethod(exportedOnly bool, name string) (Method, error) {
	structMethod, result := getStructMethodByName(t.one, name)
	if !result {
		return Method{}, errors.New(fmt.Sprintf("method named %s not found", name))
	}
	isExported := isExportedMethod(structMethod)
	if exportedOnly && !isExported {
		return Method{}, errors.New(fmt.Sprintf("method named %s is not exported", name))
	}
	return newMethod(structMethod), nil
}

func (t StructType) privateGetMethods(exportedOnly bool) ([]Method, error) {
	methods := make([]Method, 0)
	for index := 0; index < getNumMethod(t.one); index++ {
		method := getStructMethodByIndex(t.one, index)
		isExported := isExportedMethod(method)
		if exportedOnly && !isExported {
			continue
		}
		methods = append(methods, newMethod(method))
	}
	return methods, nil
}
