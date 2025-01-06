package resources

type Payee interface {
	Salary() float32
	Id() string
	FullName() string
}
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

func (s *Staff) Salary() float32 {
	return s.BaseSalary + s.Allowance
}

func (s *Staff) Id() string {
	return s.IdNumber
}

func (s *Staff) FullName() string {
	return s.Name
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

func (contractor *Contractor) Id() string {
	return contractor.IdNumber
}
func (contractor *Contractor) FullName() string {
	return contractor.Name
}
