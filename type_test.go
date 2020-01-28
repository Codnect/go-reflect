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

func TestIsString(t *testing.T) {
	typ := GetType("hello go!")
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), true)
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
}

func TestIsSlice(t *testing.T) {
	typ := GetType([]int{1, 2, 3})
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), true)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
}

func TestIsArray(t *testing.T) {
	typ := GetType([3]int{1, 2, 3})
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), true)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
}

func TestIsNumber(t *testing.T) {
	typ := GetType(3)
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), true)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
	typ = GetType(3.4)
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), true)
	assert.Equal(t, typ.IsBool(), false)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
}

func TestIsBool(t *testing.T) {
	typ := GetType(true)
	assert.Equal(t, typ.IsStruct(), false)
	assert.Equal(t, typ.IsString(), false)
	assert.Equal(t, typ.IsSlice(), false)
	assert.Equal(t, typ.IsError(), false)
	assert.Equal(t, typ.IsArray(), false)
	assert.Equal(t, typ.IsMap(), false)
	assert.Equal(t, typ.IsNumber(), false)
	assert.Equal(t, typ.IsBool(), true)
	assert.Equal(t, typ.IsFunction(), false)
	assert.Equal(t, typ.IsInterface(), false)
	assert.Equal(t, typ.IsTag(), false)
	assert.Equal(t, typ.IsPointer(), false)
}
