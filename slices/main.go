package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5, 7, 11}
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // made first slice of length 2, capacity of 4 and pointer is on "John"
	b := names[1:3] // same here slice of length 2, capacity of 3, and pointer is on "Paul"
	fmt.Println(a, b)
	
	// as slices are just reference of the original array so, if slices are modified it always be reflected on the original array as well
	b[0] = "xxx" // here slice b's first element is changed, this will reflect on the original array and on the other slices that were made by the same backing array bcz they share the same reference
	fmt.Println(a, b)
	fmt.Println(names)
}
