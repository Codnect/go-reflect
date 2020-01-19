package reflect

import (
	"github.com/codnect/go-one"
)

func GetType(one one.One) Type {
	typ := newType(one)
	return typ
}
