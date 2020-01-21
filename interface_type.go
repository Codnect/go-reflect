package reflect

import "github.com/codnect/go-one"

type InterfaceType struct {
	one one.One
}

func newInterfaceType(one one.One) InterfaceType {
	return InterfaceType{
		one: one,
	}
}
