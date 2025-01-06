package main

import (
	"fmt"
	"payroll/resources"
)

func salary(emp *resources.Employee) float32 {
	return emp.Salary()
}

func main() {
	s1 := resources.NewStaff("001", "John", 1000.0, 500.0)
	c1 := resources.NewContractor("025", "Peter", 100.0, 10)

	fmt.Printf("#%s\t%-15s\t%.2f\n", s1.IdNumber, s1.Name, s1.Salary())
	fmt.Printf("#%s\t%-15s\t%.2f\n", c1.IdNumber, c1.Name, c1.Salary())

	// fmt.Println(salary(s1)) //wont work
	// fmt.Println(salary(c1)) //wont work
}
