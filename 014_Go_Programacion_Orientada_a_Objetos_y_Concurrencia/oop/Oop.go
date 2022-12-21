package oop

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func EmployeeExample() {
	e := Employee{}
	e.id = 1
	e.name = "Ana"
	e.SetId(5)
	e.SetName("Luis")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
	e4 := NewEmployee(1, "Name 2")
	fmt.Println(e4)
}

type Person struct {
	age int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func InheritanceExample() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Maria"
	ftEmployee.age = 27
	fmt.Printf("%v", ftEmployee)
	ftEmployee.getMessage("Full time employee")
}

func (ftEmployee FullTimeEmployee) getMessage(s string) string {
	fmt.Println(s)
	return s
}
