package main

import "fmt"

// defined a struct Vertex with data x and y of both integer
type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2} // assigning both vars
	v2 = Vertex{x : 12} // assigning only one field using key : value pairs
	v3 = Vertex{} // not assigning, so by default will take the default value of int (0)
	p1 = &Vertex{22, 34} // assigning the address of a struct ; has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2}) // here 1 goes to x and 2 goes to y, just how they were defined
	// printed in string form
	v := Vertex{23, 45} // structs can be assigned to a variable 

	v.x = 32 // elements or data of any struct can be accessed using the dot notation and can be modified as well

	fmt.Println(v)

	p := &v // assigning the address of the struct to a variable

	p.x = 1e3 // modifying the value to 10^3

	fmt.Println(v, *p) // modifies the original value as well

	fmt.Println(v1, v2, v3, p1)
}
