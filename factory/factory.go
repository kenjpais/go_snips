package factory

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

type EmployeeFactory struct {
}

func NewEmployeeFactory(position string,
	annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

func TestFactory() {
	developerFactory := NewEmployeeFactory("developer", 6000)
	managerFactory := NewEmployeeFactory("manager", 9000)
	developer := developerFactory("Adam")
	manager := managerFactory("Eve")
	fmt.Println(developer, manager)
}
