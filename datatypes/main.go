package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// multiple variable declaration can be grouped here like this with different datatypes and values
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1 // this is unsigned integer
	z complex128 = cmplx.Sqrt(-5 + 12i) // this is complex and using the cmplx package's Sqrt function used to calculate the square root of the number
)

// Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.
// uses some format specifiers also called verbs like %T for type, %v for the value, the var contains, %q for strings with quotes
func main(){
	fmt.Printf("Types: %T Values: %v\n", ToBe, ToBe)
	fmt.Printf("Types: %T Values: %v\n", MaxInt, MaxInt)
	fmt.Printf("Types: %T Values: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y)) // here explicitly converting the int to float

	
	// type inference
	a := 42 // here a variable is assigned the datatype by the compiler implicitly, it can't be changed in later phases of the program
	
	// a = "string" // no allowed
	
	// constants
	
	const PI = 3.14 // using the const keyword this variable is now a constant
	
	// PI = 2.134 // not allowed as its a constant
	fmt.Println(f, a, PI)
}