package main

import (
	"fmt"
	"math"
)
// understanding anonymous functions : can be assigned to any variable for better usecase
// can be given as arguments to other functions

// here is the example of giving an anonymous function as parameter
func compute(fn func(float64, float64) float64) float64{
	return fn(3, 4)
}

func main() {
	// assigning the anonymous function to a variable
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}
	fmt.Println(hypot(5, 12)) // here calling the variable containing the function declaration using the variable name

	fmt.Println(compute(hypot)) // directly giving the variable name and default params will be given to it
	fmt.Println(compute(math.Pow)) // same goes for this one as well
}