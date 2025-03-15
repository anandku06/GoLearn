package main

import "fmt"

// example of closures in GO
// here function adder returns another function tnat takes outer function's variable and modifies it then returns
// here due to closure, the variable's state is preserved by the inner function
func adder() func(int) int {
	sum := 0
	return func (x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
	// in every iteration sum value is preserved and used for the next iteration 
}