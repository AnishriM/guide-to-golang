package main

import "fmt"

func main() {
	slice := []int{42, 28, 20, 27, 18}

	//Iterating over slice.
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i := 0; i < len(slice); {
		fmt.Println(slice[i])
		i++
	}

	println("While Loop")
	i := 0
	for {
		if i == len(slice) {
			break
		}
		fmt.Println(slice[i])
		i++

	}

	println("Iterate using range")
	for index, value := range slice {
		fmt.Println(index, " : ", value)
	}
}
