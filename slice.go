package main

import "fmt"

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	/*
		OUTPUT:
				0
				0
				[]
	*/
	fmt.Println("------------")

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	/*
		OUTPUT:
					4
					4
					[Go Slices Are Powerful]
	*/
	fmt.Println("------------")

	// Create a Slice From an Array

	var array = [6]int{1, 2, 3, 4, 5, 6}
	myslice := array[2:5]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	/*
		OUTPUT:
					myslice = [3 4 5]
					length = 3
					capacity = 4
	*/

	fmt.Println("------------")

	// Create a Slice With The make() Function

	myslicemake := make([]int, 5, 10)
	fmt.Printf("myslicemake = %v\n", myslicemake)
	fmt.Printf("length = %d\n", len(myslicemake))
	fmt.Printf("capacity = %d\n", cap(myslicemake))

	/*
				OUTPUT:
							myslicemake = [0 0 0 0 0]
		length = 5
		capacity = 10
	*/
	fmt.Println("------------")

	// with omitted capacity
	myslicemake2 := make([]int, 5)
	fmt.Printf("myslicemake = %v\n", myslicemake2)
	fmt.Printf("length = %d\n", len(myslicemake2))
	fmt.Printf("capacity = %d\n", cap(myslicemake2))
	/*
				OUTPUT:
							myslicemake = [0 0 0 0 0]
		length = 5
		capacity = 5
	*/
}
