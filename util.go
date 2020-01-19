package reflect

import (
	"errors"
	"github.com/codnect/go-one"
	"reflect"
	"unicode"
)

func isPointer(val reflect.Value) bool {
	return val.Kind() == reflect.Ptr
}

func isString(val reflect.Value) bool {
	return val.Kind() == reflect.String
}

func isStruct(val reflect.Value) bool {
	return val.Kind() == reflect.Struct
}

func getValue(one one.One) reflect.Value {
	val := reflect.ValueOf(one)
	if isPointer(val) {
		val = val.Elem()
	}
	return val
}

func getFieldValueByName(one one.One, fieldName string) (one.One, error) {
	val := getValue(one)
	if isStruct(val) {
		fieldVal := val.FieldByName(fieldName)
		if isString(fieldVal) {
			return fieldVal.String(), nil
		}
		return fieldVal.Interface(), nil
	}
	return nil, errors.New("instance must be only struct")
}

func setFieldValueByName(instance one.One, fieldName string, value one.One) error {
	val := getValue(instance)
	fieldVal := val.FieldByName(fieldName)
	if isString(fieldVal) {
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
