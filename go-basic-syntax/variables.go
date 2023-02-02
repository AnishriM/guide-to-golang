package main

import "fmt"

func main() {
	/* Way 1: Explicitly telling variable type */
	var firstName string = "Sagar"
	fmt.Println("FirstName:" + firstName)

	/* Way 2: Compiler will identify variable type */
	lastName := "Mhantati"
	fmt.Println("LastName:" + lastName)
}
