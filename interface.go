package reflect

import (
	"errors"
	"fmt"
	"github.com/codnect/go-one"
	"reflect"
)

type Interface struct {
	typ reflect.Type
}

func newInterface(p reflect.Type) Interface {
	return Interface{
		typ: p,
	}
}

func (t Interface) GetMethodByName(name string) (Method, error) {
	return t.privateGetMethod(true, name)
}

func (t Interface) GetMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t Interface) GetMethods() []Method {
	return t.privateGetMethods(true)
}

func (t Interface) GetDeclaredMethodByName(name string) (Method, error) {
	return t.privateGetMethod(false, name)
}

func (t Interface) GetDeclaredMethods() []Method {
	return t.privateGetMethods(false)
}

func (t Interface) GetDeclaredMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t Interface) privateGetMethod(exportedOnly bool, name string) (Method, error) {
	structMethod, result := getMethodByName(t.typ, name)
	if !result {
		return Method{}, errors.New(fmt.Sprintf("method named %s not found", name))
	}
	isExported := isExportedMethod(structMethod)
	if exportedOnly && !isExported {
		return Method{}, errors.New(fmt.Sprintf("method named %s is not exported", name))
	}
	return newMethod(structMethod), nil
}

func (t Interface) privateGetMethods(exportedOnly bool) []Method {
	methods := make([]Method, 0)
	for index := 0; index < getNumMethod(t.typ); index++ {
		method := getMethodByIndex(t.typ, index)
		isExported := isExportedMethod(method)
		if exportedOnly && !isExported {
			continue
		}
		methods = append(methods, newMethod(method))
	}
	return methods
}
