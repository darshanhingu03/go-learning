package main

import "fmt"

// type Constants
const A int = 1

// untype Constants
const B = 2

func main() {
	// When a constant is declared, it is not possible to change the value later
	fmt.Println(A)
	fmt.Println(B)

	fmt.Println("-------------")

	// Multiple constants can be grouped together into a block for readability:
	const (
		X int = 10
		Y     = 20
		Z     = "hello"
	)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
}
