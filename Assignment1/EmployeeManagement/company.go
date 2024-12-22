package EmployeeManagement

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-Time Employee: ID=%d, Name=%s, Salary=%d Tenge", fte.ID, fte.Name, fte.Salary)
}

func (pte PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-Time Employee: ID=%d, Name=%s, HourlyRate=%d Tenge, HoursWorked=%.2f", pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(emp Employee) {
	if c.Employees == nil {
		c.Employees = make(map[string]Employee)
	}

	switch e := emp.(type) {
	case FullTimeEmployee:
		c.Employees[fmt.Sprintf("%d", e.ID)] = emp
	case PartTimeEmployee:
		c.Employees[fmt.Sprintf("%d", e.ID)] = emp
	}
	fmt.Println("Employee added successfully!")
}

func (c Company) ListEmployees() {
	fmt.Println("Company Employees:")
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
