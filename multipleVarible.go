package main

import "fmt"

func main() {
	//  If you use the type keyword, it is only possible to declare one type of variable per line
	var a, b, c, d int = 1, 2, 3, 44
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("-------------")
	// If the type keyword is not specified, you can declare different types of variables in the same line

	var intVar, stringVar = 6, "Hello"

	fmt.Println(intVar)
	fmt.Println(stringVar)

	fmt.Println("-------------")

	intVar2, stringVar2 := 7, "hellllllo"

	fmt.Println(intVar2)
	fmt.Println(stringVar2)

	fmt.Println("-------------")

	// Multiple variable declarations can also be grouped together into a block for greater readability

	var (
		x int
		y int    = 10
		z string = "multi variable"
	)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
