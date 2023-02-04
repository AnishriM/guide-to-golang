package main

import "fmt"

func main() {
	var name string
	name = "Sagar"
	pointer := &name

	fmt.Println(name)
	fmt.Println(pointer)
	*pointer = "Anishri"
	fmt.Println(name)
}
