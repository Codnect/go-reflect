package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type FunctionType struct {
	typ reflect.Type
}

func newFunctionType(p reflect.Type) FunctionType {
	return FunctionType{
		typ: p,
	}
}

func (f FunctionType) GetReturnTypeCount() int {
	return getFunctionReturnTypeCount(f.typ)
}

func (f FunctionType) GetReturnTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getFunctionReturnTypeCount(f.typ); index++ {
		returnType := getFunctionParameterByIndex(f.typ, index)
		types = append(types, newTypeWithType(returnType))
	}
	return types
}

func (f FunctionType) GetParameterCount() int {
	return getFunctionParameterCount(f.typ)
}

func (f FunctionType) GetParameterTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getFunctionParameterCount(f.typ); index++ {
		returnType := getFunctionParameterByIndex(f.typ, index)
		types = append(types, newTypeWithType(returnType))
	}
	return types
}

func (f FunctionType) Invoke(instance one.One, params ...one.One) one.One {
	return nil
}
