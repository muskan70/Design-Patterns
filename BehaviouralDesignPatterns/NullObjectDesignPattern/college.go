package main

type College struct {
	departments []IDepartment
}

func (c *College) AddDepartment(name string, professors int) {
	if name == "Computer Science" {
		cs := &ComputerScienceDepartment{professors: professors}
		c.departments = append(c.departments, cs)
	}
	if name == "Mechanical" {
		mech := &MechanicalDepartment{professors: professors}
		c.departments = append(c.departments, mech)
	}
}

func (c *College) GetDepartment(name string) IDepartment {
	for _, d := range c.departments {
		if d.GetDepartmentName() == name {
			return d
		}
	}
	return &NullDepartment{}
}
