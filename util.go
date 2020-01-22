package reflect

import (
	"errors"
	"github.com/codnect/go-one"
	"log"
	"reflect"
	"unicode"
)

func getValue(one one.One) reflect.Value {
	val := reflect.ValueOf(one)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func getType(one one.One) reflect.Type {
	typ := reflect.TypeOf(one)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

func getFieldValueByName(one one.One, fieldName string) (one.One, error) {
	val := getValue(one)
	if val.Kind() == reflect.Struct {
		fieldVal := val.FieldByName(fieldName)
		if fieldVal.Kind() == reflect.String {
			return fieldVal.String(), nil
		}
		return fieldVal.Interface(), nil
	}
	return nil, errors.New("instance must be only struct")
}

func setFieldValueByName(instance one.One, fieldName string, value one.One) error {
	val := getValue(instance)
	fieldVal := val.FieldByName(fieldName)
	if fieldVal.Kind() == reflect.String {
		fieldVal.Set(val)
	}
	return nil
}

func isExportedField(structField reflect.StructField) bool {
	return unicode.IsUpper(rune(structField.Name[0]))
}

func isExportedMethod(method reflect.Method) bool {
	return unicode.IsUpper(rune(method.Name[0]))
}

func getStructNumField(typ reflect.Type) int {
	if typ.Kind() != reflect.Struct {
		log.Fatal("Type is not struct")
	}
	return typ.NumField()
}

func getNumMethod(typ reflect.Type) int {
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Interface {
		log.Fatal("Type is not struct or interface")
	}
	return typ.NumMethod()
}

func getStructFieldByName(typ reflect.Type, name string) (reflect.StructField, bool) {
	if typ.Kind() != reflect.Struct {
		log.Fatal("Type is not struct")
	}
	return typ.FieldByName(name)
}

func getMethodByName(typ reflect.Type, name string) (reflect.Method, bool) {
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Interface {
		log.Fatal("Type is not struct or interface")
	}
	return typ.MethodByName(name)
}

func getStructFieldByIndex(typ reflect.Type, index int) reflect.StructField {
	if typ.Kind() != reflect.Struct {
		log.Fatal("Type is not struct")
	}
	return typ.Field(index)
}

func getMethodByIndex(typ reflect.Type, index int) reflect.Method {
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Interface {
		log.Fatal("Type is not struct or interface")
	}
	return typ.Method(index)
}

func getFieldTagValueByTagName(structField reflect.StructField, tagName string) (value string, ok bool) {
	return structField.Tag.Lookup(tagName)
}

func getMethodReturnTypeCount(method reflect.Method) int {
	return method.Type.NumOut()
}

func getMethodParameterCount(method reflect.Method) int {
	return method.Type.NumIn()
}

func getMethodReturnTypeByIndex(method reflect.Method, index int) reflect.Type {
	return method.Type.Out(index)
}

func getMethodParameterByIndex(method reflect.Method, index int) reflect.Type {
	return method.Type.In(index)
}
