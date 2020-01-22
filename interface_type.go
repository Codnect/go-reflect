package reflect

import (
	"errors"
	"fmt"
	"github.com/codnect/go-one"
	"reflect"
)

type InterfaceType struct {
	typ reflect.Type
}

func newInterfaceType(p reflect.Type) InterfaceType {
	return InterfaceType{
		typ: p,
	}
}

func (t InterfaceType) GetMethodByName(name string) (Method, error) {
	return t.privateGetMethod(true, name)
}

func (t InterfaceType) GetMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t InterfaceType) GetMethods() []Method {
	return t.privateGetMethods(true)
}

func (t InterfaceType) GetDeclaredMethodByName(name string) (Method, error) {
	return t.privateGetMethod(false, name)
}

func (t InterfaceType) GetDeclaredMethods() []Method {
	return t.privateGetMethods(false)
}

func (t InterfaceType) GetDeclaredMethod(parameters ...one.One) (Method, error) {
	return Method{}, nil
}

func (t InterfaceType) privateGetMethod(exportedOnly bool, name string) (Method, error) {
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

func (t InterfaceType) privateGetMethods(exportedOnly bool) []Method {
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
