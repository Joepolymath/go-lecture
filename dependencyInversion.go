package main

import "fmt"

// This principle states that high-level modules should not depend on low-level modules, but rather both should depend on abstractions. This helps to reduce the coupling between components and make the code more flexible and maintainable.

type Worker struct {
	ID		uint64
	Name	string
	Role	string
}

type Supervisor struct {
	ID		uint64
	Name	string
	Role	string
}

type Employee interface {
	GetID() 		uint64
	GetName()	string
	GetRole()	string
}

// High level module not relying directly on low level module. But abstraction.
type Department struct { 
	Employees	[]Employee
	Name			string
}

// Get IDs
func (e Worker) GetID() uint64 {
	return e.ID
}
func (e Supervisor) GetID() uint64 {
	return e.ID
}

// Get Names
func (e Worker) GetName() string {
	return e.Name
}
func (e Supervisor) GetName() string {
	return e.Name
}

// Get Roles
func (e Worker) GetRole() string {
	return e.Role
}
func (e Supervisor) GetRole() string {
	return e.Role
}

func (d *Department) GetEmployee(name string) (res Employee) {
	for _, employee := range d.Employees {
		if employee.GetName() == name {
			res = employee
		}
	}
	return
}

func (d *Department) AddEmployee(employee Employee) (res []Employee) {
	d.Employees = append(d.Employees, employee)
	res = d.Employees
	return
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, employee := range d.Employees {
		res = append(res, employee.GetName())
	}
	return
}

func DIP() {
	d := &Department{}
	d.AddEmployee(Worker{
		Name: "Joshua",
		Role: "Backend Engineer",
		ID: 312,
	})
	d.AddEmployee(Supervisor{
		Name: "Dozie",
		Role: "Engineering Manager",
		ID: 12,
	})
	
	fmt.Println(d.GetEmployeeNames())
}