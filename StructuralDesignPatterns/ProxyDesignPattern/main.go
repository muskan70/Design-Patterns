package main

func main() {
	employeeDB := NewEmployeeProxy()
	employeeDB.Create("ADMIN", "Mayank")
	employeeDB.Create("ADMIN", "Sharan")
	employeeDB.Create("MANAGER", "Karan")

	employeeDB.Get("ADMIN", 1)
	employeeDB.Get("MANAGER", 2)
	employeeDB.Get("ADMIN", 2)
	employeeDB.Get("ADMIN", 3)
	employeeDB.Get("MANAGER", 3)

	employeeDB.Delete("ADMIN", 1)
	employeeDB.Delete("MANAGER", 2)

}
