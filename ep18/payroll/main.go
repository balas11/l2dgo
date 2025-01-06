package main

import (
	"fmt"
	"payroll/resources"
)

func salary(payee resources.Payee) float32 {
	return payee.Salary()
}

func main() {
	payees := []resources.Payee{
		resources.NewStaff("001", "John", 1000.0, 500.0),
		resources.NewContractor("025", "Peter", 100.0, 10),
	}

	for _, p := range payees {
		fmt.Printf("#%s\t%-15s\t%.2f\n", p.Id(), p.FullName(), salary(p))
	}

	salary(payees[0])
	salary(payees[1])
}
