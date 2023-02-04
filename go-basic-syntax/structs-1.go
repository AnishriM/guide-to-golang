package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

func main() {
	//Way1
	engineer1 := Engineer{Name: "Sagar", Age: 25}
	fmt.Println(engineer1)

	//Way2
	engineer2 := Engineer{"Anishri", 25}
	fmt.Println(engineer2)
}
