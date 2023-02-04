package main

import "fmt"

// Function without argument and return type.
func HelloWorld() {
	println("Hello World")
}

// Function with argument and without return type.
func HelloPerson(person string) {
	fmt.Println("Hello", person)
}

// Function with argument and return type.
func Total(a int, b int) int {
	return a + b
}

// Function with argument and multiple return type.
func SumAndMultiplication(a int, b int) (int, int) {
	return a + b, a * b
}
func main() {
	HelloWorld()
	HelloPerson("Sagar")
	total := Total(1, 2)
	fmt.Println("Total:", total)

	sum, multiplication := SumAndMultiplication(2, 3)
	fmt.Println("Sum:", sum, "Multiplication:", multiplication)
}
