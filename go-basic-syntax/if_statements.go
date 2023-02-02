package main

func main() {
	println("If statments in golang")

	var height int = 130
	age := 16

	if height >= 180 || age >= 18 {
		println("Customer can access the rides.")
	} else if height >= 140 && age >= 15 {
		println("Customer can access childern's rides.")
	} else {
		println("Customer cannont access any rides.")
	}
}
