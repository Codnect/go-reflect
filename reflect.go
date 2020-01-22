package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

func GetType(one one.One) Type {
	typ := newType(reflect.TypeOf(one), reflect.ValueOf(one))
	return typ
}
