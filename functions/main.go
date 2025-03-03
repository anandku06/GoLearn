package main

import "fmt"

// if both params have same datatype then (x, y int) works
func add(x int, y int) int {
	return x + y
}
// here add is a custom function made that takes two parameter of integers and returns an integer as well
// GO is a staically-typed lang => datatype checking is done during compilation

// function with multiple returns
// here swap func takes two params and returns two strings so used parenthesis
func swap(x, y string) (string, string){
	return y, x
}

// here a new type of return method
// in place of giving (int, int), we used (x, y int) i.e. used the return variables and thier datatypes
// what this does that now x and y are initialised before the start of the function body's execution
// that's why x and y are not initialised as well
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} // not recommended ; causes readabilty issues

func main(){
	fmt.Println(add(43,45)) // here the function is called with the right params
	// fmt.Println(add("43",45)) // if the datatype of the args are diff, then throws error
	a, b := swap("world", "hello") // using the comma-seperated variables to assign the two returns from the function 

	fmt.Println(a, b)

	fmt.Println(split(34))
}