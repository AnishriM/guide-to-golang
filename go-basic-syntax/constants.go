package main

import "fmt"

func main() {
	//typed constants
	const firstName string = "Sagar"
	fmt.Println("FirstName: " + firstName)

	//Untyped constants
	const lastName = "Mhantati"
	fmt.Println("LastName:" + lastName)

	/* 	shorthand code is not supported while declaring constant
	e.g lastName := "Mhantati"
	*/
}
