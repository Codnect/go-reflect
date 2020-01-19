package reflect

import (
	"errors"
	"github.com/codnect/go-one"
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

func getNumField(one one.One) int {
	return getValue(one).NumField()
}

func getStructFieldByName(one one.One, name string) (reflect.StructField, bool) {
	return getValue(one).Type().FieldByName(name)
}

func getStructFieldByIndex(one one.One, index int) reflect.StructField {
	return getValue(one).Type().Field(index)
}
