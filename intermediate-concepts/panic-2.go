package main

import "fmt"

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName() {
	e.Name = "Anishri"
}

func doStuff(e *Engineer) {
	defer e.UpdateName()
	println("Doing exciting stuff.")
}
func main() {
	engineer := Engineer{
		Name: "Sagar",
	}
	fmt.Println(engineer)
	doStuff(&engineer)
	fmt.Println(engineer)
}

/*
Output:
@AnishriM ➜ /workspaces/guide-to-golang/intermediate-concepts (go-basic-syntax) $ go run panic-2.go 
{Sagar}
Doing exciting stuff.
{Anishri}
*/