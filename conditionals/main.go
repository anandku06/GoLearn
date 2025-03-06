package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// use of if keyword with the condition not inside any parenthesis
	// if the condition is true then, execute the code inside the if-body or execute the else-body
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// another type of if-else case you may see
	// initialisation with condition checking, but the scope of the initialisation is in this block only
	if v := math.Pow(x, n); v < lim {
		return v
	}else{
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// fmt.Println(v) // scope of the v variable is only to the if-else block ; throws error
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10))
}
