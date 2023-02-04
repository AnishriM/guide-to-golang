package main

import "fmt"

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

func doStuff(e *Engineer) {
	name := "Anishri"
	defer e.UpdateName(name)
	println("Doing exciting stuff.")
	name = "Shiva"
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
@AnishriM ➜ /workspaces/guide-to-golang/intermediate-concepts (go-basic-syntax) $ go run panic-3.go
{Sagar}
Doing exciting stuff.
{Anishri}

Note:
It will not update name with Shiva because it evaluetes the paratmeter of defer at the time of its first occurance.
*/
