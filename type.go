package reflect

import "github.com/codnect/go-one"

type Type struct {
	name string
}

func newType(one one.One) *Type {
	return &Type{}
}

func (t *Type) GetName() string {
	return t.name
}
