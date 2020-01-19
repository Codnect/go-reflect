package reflect

import (
	"errors"
	"fmt"
	"github.com/codnect/go-one"
	"reflect"
	"strconv"
)

type Field struct {
	structField reflect.StructField
	BaseTaggedElement
}

func newField(structField reflect.StructField) Field {
	return Field{
		structField: structField,
	}
}

func (field Field) GetName() string {
	return field.structField.Name
}

func (field Field) GetDeclaringType() *Type {
	return nil
}

func (field Field) GetModifiers() int {
	return 0
}

func (field Field) IsPrimitive() bool {
	return false
}

func (field Field) IsStruct() bool {
	return false
}
func (field Field) IsAnonymous() bool {
	return field.structField.Anonymous
}

func (field Field) IsExported() bool {
	return isExportedField(field.structField)
}

func (field Field) Get(instance one.One) (one.One, error) {
	return getFieldValueByName(instance, field.structField.Name)
}

func (field Field) GetBool(instance one.One) (bool, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(bool), err
}

func (field Field) GetByte(instance one.One) (byte, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(byte), err
}

func (field Field) GetInt8(instance one.One) (int8, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int8), err
}

func (field Field) GetShort(instance one.One) (int16, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int16), err
}

func (field Field) GetInt16(instance one.One) (int16, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int16), err
}

func (field Field) GetInt(instance one.One) (int, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int), err
}

func (field Field) GetInt32(instance one.One) (int32, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int32), err
}

func (field Field) GetInt64(instance one.One) (int64, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(int64), err
}

func (field Field) GetString(instance one.One) (string, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(string), err
}

func (field Field) GetFloat32(instance one.One) (float32, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(float32), err
}

func (field Field) GetFloat64(instance one.One) (float64, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(float64), err
}

func (field Field) GetUint8(instance one.One) (uint8, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(uint8), err
}

func (field Field) GetUInt16(instance one.One) (uint16, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(uint16), err
}

func (field Field) GetUInt(instance one.One) (uint32, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(uint32), err
}

func (field Field) GetUInt32(instance one.One) (uint64, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(uint64), err
}

func (field Field) GetUInt64(instance one.One) (uint64, error) {
	result, err := getFieldValueByName(instance, field.structField.Name)
	return result.(uint64), err
}

func (field Field) Set(instance one.One, value one.One) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetBool(instance one.One, value bool) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetByte(instance one.One, value byte) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetInt8(instance one.One, value int8) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetShort(instance one.One, value int16) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetInt16(instance one.One, value int16) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetInt(instance one.One, value int) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetInt32(instance one.One, value int32) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetInt64(instance one.One, value int64) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetString(instance one.One, value string) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetFloat32(instance one.One, value float32) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetFloat64(instance one.One, value float64) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetUint8(instance one.One, value uint8) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetUInt16(instance one.One, value uint16) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetUInt(instance one.One, value uint) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetUInt32(instance one.One, value uint32) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) SetUInt64(instance one.One, value uint64) error {
	return setFieldValueByName(instance, field.structField.Name, value)
}

func (field Field) GetTag(tag Tag) (Tag, error) {
	return field.GetTagByName(tag.GetTagName())
}

func (field Field) GetTags() []Tag {
	fieldTags := make([]Tag, 0)
	tag := field.structField.Tag
	for tag != "" {
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		quotedValue := string(tag[:i+1])
		tag = tag[i+1:]

		value, err := strconv.Unquote(quotedValue)
		if err != nil {
			break
		}

		fieldTag := newFieldTag(name, value)
		fieldTags = append(fieldTags, fieldTag)
	}
	return fieldTags
}

func (field Field) GetTagByName(name string) (Tag, error) {
	value, ok := getFieldTagValueByTagName(field.structField, name)
	if ok {
		tag := newFieldTag(name, value)
		return tag, nil
	}
	errText := fmt.Sprintf("Tag named %s not found ", name)
	return FieldTag{}, errors.New(errText)
}

func (field Field) Equals(other one.One) bool {
	return false
}

func (field Field) HashCode() int {
	return 0
}

func (field Field) String() string {
	return ""
}
