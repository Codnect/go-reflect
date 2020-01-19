package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type Field struct {
	name string
}

func newField(name string) Field {
	return Field{
		name: name,
	}
}

func (field Field) GetName() string {
	return field.name
}

func (field Field) IsPrimitive() bool {
	return true
}

func (field Field) IsStruct() bool {
	return true
}

func (field Field) Get(instance one.One) (one.One, error) {
	return getFieldValueByName(instance, field.name)
}

func (field Field) GetBool(instance one.One) (bool, error) {
	result, err := field.Get(instance)
	return result.(bool), err
}

func (field Field) GetByte(instance one.One) (byte, error) {
	result, err := field.Get(instance)
	return result.(byte), err
}

func (field Field) GetInt8(instance one.One) (int8, error) {
	result, err := field.Get(instance)
	return result.(int8), err
}

func (field Field) GetShort(instance one.One) (int16, error) {
	result, err := field.Get(instance)
	return result.(int16), err
}

func (field Field) GetInt16(instance one.One) (int16, error) {
	result, err := field.Get(instance)
	return result.(int16), err
}

func (field Field) GetInt(instance one.One) (int, error) {
	result, err := field.Get(instance)
	return result.(int), err
}

func (field Field) GetInt32(instance one.One) (int32, error) {
	result, err := field.Get(instance)
	return result.(int32), err
}

func (field Field) GetInt64(instance one.One) (int64, error) {
	result, err := field.Get(instance)
	return result.(int64), err
}

func (field Field) GetString(instance one.One) (string, error) {
	result, err := field.Get(instance)
	return result.(string), err
}

func (field Field) GetFloat32(instance one.One) (float32, error) {
	result, err := field.Get(instance)
	return result.(float32), err
}

func (field Field) GetFloat64(instance one.One) (float64, error) {
	result, err := field.Get(instance)
	return result.(float64), err
}

func (field Field) GetUint8(instance one.One) (uint8, error) {
	result, err := field.Get(instance)
	return result.(uint8), err
}

func (field Field) GetUInt16(instance one.One) (uint16, error) {
	result, err := field.Get(instance)
	return result.(uint16), err
}

func (field Field) GetUInt(instance one.One) (uint32, error) {
	result, err := field.Get(instance)
	return result.(uint32), err
}

func (field Field) GetUInt32(instance one.One) (uint64, error) {
	result, err := field.Get(instance)
	return result.(uint64), err
}

func (field Field) GetUInt64(instance one.One) (uint64, error) {
	result, err := field.Get(instance)
	return result.(uint64), err
}

func (field Field) Set(instance one.One, value one.One) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetBool(instance one.One, value bool) {
	getValue(instance).FieldByName(field.name).SetBool(value)
}

func (field Field) SetByte(instance one.One, value byte) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetInt8(instance one.One, value int8) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetShort(instance one.One, value int16) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetInt16(instance one.One, value int16) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetInt(instance one.One, value int) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetInt32(instance one.One, value int32) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetInt64(instance one.One, value int64) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetString(instance one.One, value string) {
	getValue(instance).FieldByName(field.name).SetString(value)
}

func (field Field) SetFloat32(instance one.One, value float32) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetFloat64(instance one.One, value float64) {
	getValue(instance).FieldByName(field.name).SetFloat(value)
}

func (field Field) SetUint8(instance one.One, value uint8) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetUInt16(instance one.One, value uint16) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetUInt(instance one.One, value uint) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetUInt32(instance one.One, value uint32) {
	getValue(instance).FieldByName(field.name).Set(reflect.ValueOf(value))
}

func (field Field) SetUInt64(instance one.One, value uint64) {
	getValue(instance).FieldByName(field.name).SetUint(value)
}
