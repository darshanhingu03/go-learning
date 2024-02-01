package main

import "fmt"

var outsideFuncation string

func main() {
	var stringVar string = "darshan" //type is string
	var integerVar = 2               //type is inferred
	x := "hingu"                     //type is inferred

	fmt.Println(stringVar)
	fmt.Println(integerVar)
	fmt.Println(x)

	fmt.Println("-------------")
	// if you declare a variable without an initial value, its value will be set to the default value of its type:
	var a string
	var b int
	var c bool

	fmt.Println(a) // return blank string
	fmt.Println(b) // return 0
	fmt.Println(c) // return false

	fmt.Println("-------------")
	// value assign after declaration

	a = "after assign"

	fmt.Println(a)

	fmt.Println("-------------")
	// decalare outside of function

	outsideFuncation = "outside"

	fmt.Println(outsideFuncation)

}
