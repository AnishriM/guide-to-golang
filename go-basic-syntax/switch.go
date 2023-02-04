package main

import "runtime"

func main() {
	println("Switch statement in golang")
	var height int = 140
	age := 16

	/* Switch statment without expression */
	switch {
	case height >= 180 && age >= 18:
		println("Can access rides.")
	case height <= 150 && age >= 15:
		println("Can access childern's ride.")
	default:
		println("Cannot access any rides")
	}

	/* Switch statment with expression */
	switch os := runtime.GOOS; os {
	case "linux":
		println("Linux system")
	case "windows":
		println("Windows system")
	default:
		println("System not supported")
	}
}
