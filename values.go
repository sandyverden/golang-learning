package main

import "fmt"

func main() {
	// string, with can be added together with +
	fmt.Println("go" + "lang")

	// integer and float too
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Booleans, with boolean operator as you'd expect
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
