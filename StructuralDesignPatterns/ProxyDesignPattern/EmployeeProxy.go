package main

import "fmt"

type EmployeeProxy struct {
	emp                Employee
	maxAllowedRequests int
	rateLimiter        map[string]int
}

func NewEmployeeProxy() IEmployee {
	return &EmployeeProxy{
		emp:                Employee{},
		maxAllowedRequests: 2,
		rateLimiter:        make(map[string]int)}
}

func (e *EmployeeProxy) Create(user string, name string) {
	if user != "ADMIN" {
		fmt.Println("This user is not allowed to create employee")
		return
	}
	e.emp.Create(user, name)
}

func (e *EmployeeProxy) Delete(user string, id int) {
	if user != "ADMIN" {
		fmt.Println("This user is not allowed to delete employee")
		return
	}
	e.emp.Delete(user, id)
}

func (e *EmployeeProxy) Get(user string, id int) {
	times, ok := e.rateLimiter[user]
	if ok && times == e.maxAllowedRequests {
		fmt.Println("This user has exceeded max limit to access employees data")
		return
	}
	e.emp.Get(user, id)
	e.rateLimiter[user]++
}
