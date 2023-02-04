package main

import "fmt"

func main() {
	println("Array and Slices.")

	/*-----------------------Arrays-------------------------------------------------*/
	//Way 1
	students := [5]string{"Sagar", "Rob", "Alex", "Rohshan", "Rohan"}
	fmt.Println(students)

	//Way2
	var cities [3]string
	cities[0] = "Solapur"
	cities[1] = "Pune"
	cities[2] = "Hyderabad"
	fmt.Println(cities)

	/*-----------------------Slices-------------------------------------------------*/
	//Way 1
	studentsslice := []string{"Sagar", "Rob", "Alex", "Roshan", "Rohan"}
	fmt.Println(studentsslice)

	//Way 2
	var citySlice []string
	citySlice = append(citySlice, "Solapur")
	citySlice = append(citySlice, "Pune")
	citySlice = append(citySlice, "Hyderabad")
	fmt.Println(citySlice)
	fmt.Println("size of slice: ", len(citySlice))

}
