package resources

type Employee struct {
	IdNumber   string
	Name       string
	BaseSalary float32
}

func (emp *Employee) Salary() float32 {
	return emp.BaseSalary
}

type Staff struct {
	Employee
	Allowance float32
}

func NewStaff(id, name string, base, allowance float32) *Staff {
	return &Staff{
		Employee: Employee{
			IdNumber:   id,
			Name:       name,
			BaseSalary: base,
		},
		Allowance: allowance,
	}
}

func (staff *Staff) Salary() float32 {
	return staff.BaseSalary + staff.Allowance
}

type Contractor struct {
	Employee
	Rate  float32
	Hours int
}

func NewContractor(id, name string, rate float32, hours int) *Contractor {
	return &Contractor{
		Employee: Employee{
			IdNumber:   id,
			Name:       name,
			BaseSalary: 0.0,
		},
		Rate:  rate,
		Hours: hours,
	}
}

func (contractor *Contractor) Salary() float32 {
	return float32(contractor.Hours) * contractor.Rate
}
