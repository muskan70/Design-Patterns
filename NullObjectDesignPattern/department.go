package main

type IDepartment interface {
	GetNoOfProfessors() int
	GetDepartmentName() string
}

type ComputerScienceDepartment struct {
	professors int
}

func (c *ComputerScienceDepartment) GetNoOfProfessors() int {
	return c.professors
}

func (c *ComputerScienceDepartment) GetDepartmentName() string {
	return "Computer Science"
}

type MechanicalDepartment struct {
	professors int
}

func (m *MechanicalDepartment) GetNoOfProfessors() int {
	return m.professors
}

func (m *MechanicalDepartment) GetDepartmentName() string {
	return "Mechanical"
}

type NullDepartment struct {
	professors int
}

func (c *NullDepartment) GetNoOfProfessors() int {
	return 0
}

func (c *NullDepartment) GetDepartmentName() string {
	return "null Department"
}
