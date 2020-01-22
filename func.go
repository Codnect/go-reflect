package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type Function struct {
	typ reflect.Type
}

func newFunction(p reflect.Type) Function {
	return Function{
		typ: p,
	}
}

func (f Function) GetReturnTypeCount() int {
	return getFunctionReturnTypeCount(f.typ)
}

func (f Function) GetReturnTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getFunctionReturnTypeCount(f.typ); index++ {
		returnType := getFunctionParameterByIndex(f.typ, index)
		types = append(types, newType(returnType))
	}
	return types
}

func (f Function) GetParameterCount() int {
	return getFunctionParameterCount(f.typ)
}

func (f Function) GetParameterTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getFunctionParameterCount(f.typ); index++ {
		returnType := getFunctionParameterByIndex(f.typ, index)
		types = append(types, newType(returnType))
	}
	return types
}

func (f Function) Invoke(instance one.One, params ...one.One) one.One {
	return nil
}
