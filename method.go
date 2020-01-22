package reflect

import (
	"github.com/codnect/go-one"
	"reflect"
)

type Method struct {
	method reflect.Method
}

func newMethod(method reflect.Method) Method {
	return Method{
		method: method,
	}
}

func (t Method) GetReturnTypeCount() int {
	return getMethodReturnTypeCount(t.method)
}

func (t Method) GetReturnTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getMethodReturnTypeCount(t.method); index++ {
		returnType := getMethodReturnTypeByIndex(t.method, index)
		types = append(types, newTypeWithType(returnType))
	}
	return types
}

func (t Method) GetParameterCount() int {
	return getMethodParameterCount(t.method)
}

func (t Method) GetParameterTypes() []Type {
	types := make([]Type, 0)
	for index := 0; index < getMethodParameterCount(t.method); index++ {
		returnType := getMethodParameterByIndex(t.method, index)
		types = append(types, newTypeWithType(returnType))
	}
	return types
}

func (t Method) Invoke(instance one.One, params ...one.One) one.One {
	return nil
}
