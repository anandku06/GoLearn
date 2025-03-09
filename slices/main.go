package main

import (
	"fmt"
)

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

	// without any backing array
	q := []int{2, 3, 4, 5, 56} // here no backing array is involved and slice because no size is given inside the square braces
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	println(r)

	newS := []struct {
		i int
		b bool
	}{
		{2, true},
		{4, false},
		{3, true},
		{5, true},
		{21, false},
		{12, true},
	} // slice literal initialisation for struct
	fmt.Println(newS)

	// slices can also be more slices
	VeryNewS := []int{2, 3, 5, 7, 11, 13}

	VeryNewS = VeryNewS[1:4]
	fmt.Println(VeryNewS)

	VeryNewS = VeryNewS[:2] // means slicing from first index to 2
	fmt.Println(VeryNewS)

	VeryNewS = VeryNewS[1:] // means slicing from 1 to last index
	fmt.Println(VeryNewS)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0] // giving it zero length, starting from the 0 and ending to 0, as last index is excluded so empty slice
	printSlices(s)

	s = s[:4] // extending its length
	printSlices(s)

	s = s[2:] // drop its first two elements
	printSlices(s)

	var d []int // here slice hasn't got any value so it will get it's default value
	fmt.Println(d, len(d), cap(d))

	if d == nil {
		fmt.Println("nil")
	}
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// The len built-in function returns the length of v, according to its type
// The cap built-in function returns the capacity of v, according to its type:
