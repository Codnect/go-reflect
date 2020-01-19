package reflect

import (
	"github.com/codnect/go-one"
)

func GetType(one one.One) Type {
	typ := Type{}
	fields := getFields(one)
	typ.setFields(fields)
	return typ
}

func getFields(one one.One) []Field {
	val := getValue(one)
	fields := make([]Field, 0)
	for index := 0; index < val.NumField(); index++ {
		structField := val.Type().Field(index)
		field := newField(structField.Name)
		fields = append(fields, field)
	}
	return fields
}
