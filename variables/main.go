package main

import "fmt"

// vars can be initialised outside and inside of the main function as well
// outside the main function only use var for variable declaration
var c, python, java bool // variables are of same type so grouped them
// i := 20 // not allowed 

func main(){
	// var i int // "i" is initialised to its default value i.e. 0
	i := 10 // another way to initialise the variables but inside the function ONLY
	var j = 12 // can be done like this as well

	var mssg string = "Hello World"
	fmt.Println(i, c, python, java, j, mssg)
}