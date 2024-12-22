package main

import (
	"github.com/Eternality777/Nursultan_Serikuly/Assignment1/EmployeeManagement"
)

func main() {
	company := EmployeeManagement.Company{}

	ftEmp := EmployeeManagement.FullTimeEmployee{
		ID:     1,
		Name:   "Nursultan Serikuly",
		Salary: 500000,
	}

	ptEmp := EmployeeManagement.PartTimeEmployee{
		ID:          2,
		Name:        "Serik Nursultanuly",
		HourlyRate:  5000,
		HoursWorked: 25.5,
	}

	company.AddEmployee(ftEmp)
	company.AddEmployee(ptEmp)

	company.ListEmployees()
}
