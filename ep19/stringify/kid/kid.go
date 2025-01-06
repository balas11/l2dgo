package kid

import "fmt"

type Kid struct {
	name string
	age  int8
}

func New(name string, age int8) *Kid {
	return &Kid{
		name: name,
		age:  age,
	}
}

func (k *Kid) Name() string {
	return k.name
}

func (k *Kid) Age() int8 {
	return k.age
}

func (k *Kid) String() string {
	return fmt.Sprintf("%s is %d years old", k.name, k.age)
}
