package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// traditional for loop
	for i := 0; i < len(pow); i++ {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}

	// for loop with range
	// ranges over the whole iterable and assigning is done to index and value
	// i for the index and v for the value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// can give underscore if we don't want that value to be assigned as this function returns two value but need is of one
}
