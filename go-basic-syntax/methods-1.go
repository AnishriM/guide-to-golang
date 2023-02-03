package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() {
	println("Engineer's Information:")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Project Name: %s\n", e.Project.Name)
}
func main() {
	engineer := Engineer{
		Name: "Sagar",
		Age:  25,
		Project: Project{
			Name:         "Begineer's guide for go",
			Priority:     "High",
			Technologies: []string{"GO", "HTML", "CSS"},
		},
	}
	engineer.Print()
}
