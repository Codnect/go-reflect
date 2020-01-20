package reflect

import (
	"log"
	"testing"
)

type Employee struct {
	name   string `    d:""     pam:""`
	Age    int
	Salary int
}

func (emp Employee) GetAge() int {
	return emp.Age
}

func (emp Employee) GetSalary() int {
	return emp.Salary
}

type CustomTag struct {
	name   string `d:""     pam:""`
	Age    int
	Salary int
}

func (c CustomTag) GetTagName() string {
	return ""
}

func TestGetType(t *testing.T) {
	y := GetType(Employee{})
	y.GetDeclaredFields()
	z := y.NewInstancePointer()
	log.Println(z)
	instance := Employee{name: "hello", Age: 10}
	typ := GetType(instance)
	log.Println(typ.IsTag())
	x := typ.GetDeclaredFields()
	d := x[0].IsTagPresent(&Json{})
	log.Println(d)
	json, _ := x[0].GetTag(&Json{})
	log.Println(json)
	log.Println(typ.IsPointer())
	log.Println(typ.IsStruct())
	log.Println(typ.IsInterface())
	typ2 := GetType(&CustomTag{})
	log.Println(typ2.IsTag())
}
