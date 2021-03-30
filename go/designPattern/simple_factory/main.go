package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

const (
	FrontDeveloper = iota + 1
	BackendDeveloper
)

func NewEmployee(role int) *Employee {
	switch role {
	case FrontDeveloper:
		return &Employee{"", "FrontDeveloper", 60000}
	case BackendDeveloper:
		return &Employee{"", "BackendDeveloper", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(FrontDeveloper)
	m.Name = "Yo"
	fmt.Println(m)
}
