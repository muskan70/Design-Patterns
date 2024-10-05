package main

import "fmt"

func main() {
	college1 := College{}
	college1.AddDepartment("Computer Science", 5)
	college1.AddDepartment("Mechanical", 2)

	college2 := College{}
	college2.AddDepartment("Computer Science", 2)

	departments := []string{"Computer Science", "Mechanical", "Civil", "Electrical"}

	for _, dName := range departments {
		d := college1.GetDepartment(dName)
		fmt.Println("College1 :", dName, ",", d.GetNoOfProfessors())

		d = college2.GetDepartment(dName)
		fmt.Println("College2 :", dName, ",", d.GetNoOfProfessors())
	}
}
