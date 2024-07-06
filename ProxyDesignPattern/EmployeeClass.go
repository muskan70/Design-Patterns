package main

import "fmt"

var i = 1
var employeeList = make(map[int]string)

type Employee struct {
}

func (e *Employee) Create(user string, name string) {
	employeeList[i] = name
	fmt.Println("Employee Created Successfully with Id", i)
	i++
}

func (e *Employee) Delete(user string, id int) {
	delete(employeeList, id)
	fmt.Println("Employee Deleted Successfully with Id", id)
}

func (e *Employee) Get(user string, id int) {
	name, ok := employeeList[id]
	if ok {
		fmt.Println("Employee successfully extracted", id, name)
	} else {
		fmt.Println("Employee doesn't exist with id", id)
	}
}
