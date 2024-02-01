package main

import (
	"fmt"
)

func main() {
	// The Print() function prints its arguments with their default format

	var i, j string = "Hello", "World"
	fmt.Print(i)
	fmt.Print(j)
	// output: HelloWorld
	fmt.Println("\n-------------")

	// If we want to print the arguments in new lines, we need to use \n.
	fmt.Print(i, "\n", j)

	/*
	   output :	Hello
	   		   	World
	*/
	fmt.Println("\n-------------")

	// If we want to add a space between string arguments, we need to use " "
	fmt.Print(i, " ", j)

	// output: Hello World
	fmt.Println("\n-------------")

	// Print() inserts a space between the arguments if neither are strings
	x, y := 10, 20
	fmt.Print(x, y)

	// output: 10 20
	fmt.Println("\n-------------")

	// The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end

	fmt.Println(i, j)
	// output: Hello World

	// The Printf() function first formats its argument based on the given formatting verb and then prints them

	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("i has value: %v and type: %T\n", x, x)

	/*
		output: i has value: Hello and type: string
				i has value: 10 and type: int
	*/
}
