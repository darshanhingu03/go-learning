package main

import "fmt"

func main() {
	// fixed length array
	var arr1 = [3]int{1, 2, 3}

	fmt.Println(arr1)
	//OUTPUT : [1 2 3]

	/*
		var arr1 = [3]int{1, 2, 3,4}
		output: index 3 is out of bounds (>= 3)
	*/

	fmt.Println("------------")

	// inferred lengths
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	// OUTPUT: [1 2 3 4 5]

	fmt.Println("------------")

	// array of string
	arr3 := [3]string{"hello", "hiii!", "welcome"}
	fmt.Println(arr3)
	// OUTPUT: [hello hiii! welcome]

	fmt.Println("------------")

	// Access Elements of an Array
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])

	/* OUTPUT: 	1
	2
	*/

	fmt.Println("------------")

	// Change Elements of an Array

	arr2[1] = 20
	fmt.Println(arr2)
	// OUTPUT: [1 20 3 4 5]

	fmt.Println("------------")

	// Array Initialization

	arr4 := [5]int{}              //not initialized
	arr5 := [5]int{1, 2}          //partially initialized
	arr6 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)

	/*
		OUTPUT:
				[0 0 0 0 0]
				[1 2 0 0 0]
				[1 2 3 4 5]

	*/
	fmt.Println("------------")

	// Initialize Only Specific Elements

	arr7 := [5]int{1: 30, 4: 10, 2: 22}
	fmt.Println(arr7)

	// OUTPUT: [0 30 22 0 10]

	fmt.Println("------------")

	// Find the Length of an Array

	fmt.Println(len(arr7))

	// OUTPUT: 5
}
