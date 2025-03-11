package main

import "fmt"

// append() -> appends an element to the last of a slice
// takes two args : one is the slice and second is the element to be appneded; any number of elements bcz its a variadic function
func main(){
	var s []int // declared an empty slice
	printSlices(s)
	
	s = append(s, 0) // appended 0 to the end and returns
	printSlices(s)

	s = append(s, 1) // same goes for here as well
	printSlices(s)

	s = append(s, 2, 3, 4) // here multiple elements are appended together
	printSlices(s)

	var nums []int
	for i := 0; i < 20; i++ {
		nums = append(nums, i)
		printSlices(nums)
	}
	// whenever the capacity is exhausted, in the next append the caoacity is doubled
	// these extra calculations are prevented using the make function, for assigning any slice
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// append() function BTS does a number of things: first checks the given slice's length and capacity, 
// IF the length and capacity is enough for the new element, then it just appends 
// ELSE it assigns another slice with the enough length and capacity and then after appending the element it returns the new slice not the previous one