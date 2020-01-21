package reflect

import "reflect"

type Method struct {
	method reflect.Method
}

func newMethod(method reflect.Method) Method {
	return Method{
		method: method,
	}
}
