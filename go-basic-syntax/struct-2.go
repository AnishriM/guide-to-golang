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

func main() {
	engineer := Engineer{
		Name: "Sagar",
		Age:  25,
		Project: Project{
			Name:         "Beginner's guide to go",
			Priority:     "High",
			Technologies: []string{"Go", "HTML", "CSS"},
		},
	}
	fmt.Println(engineer)
	fmt.Println("Engineer Name:", engineer.Name, "\nProject Name:", engineer.Project.Name)
}
