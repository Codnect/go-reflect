package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ClickListener interface {
	OnClick()
	OnLongClick()
}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float32
	Address   string
}

func TestIsStruct(t *testing.T) {
	typ := GetType(Employee{})
	assert.Equal(t, typ.IsStruct(), true)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
	typ = GetType(&Employee{})
	assert.Equal(t, typ.IsStruct(), true)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), true)
}

func TestIsInterface(t *testing.T) {
	typ := GetType((*ClickListener)(nil))
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), true)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), true)
}
