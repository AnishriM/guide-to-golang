package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func main() {
	engineer := Engineer{Name: "Sagar"}
	fmt.Println(engineer.GetName())

	manager := Manager{Name: "XYZ"}
	fmt.Println(manager.GetName())
}
