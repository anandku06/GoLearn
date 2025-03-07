package main

import "fmt"

// defer -> used to run the program parallely with the main function
// the code with this keyword stacked somewhere and after the main function execution, it follows the LIFO and executes the defer statement

func main(){
	defer fmt.Println("World!!") // this line of code will executed after the main function is executed
	// this will executed last
	fmt.Println("Hello, ")

	fmt.Println("Counting...")

	for i := 1; i<= 5; i++ {
		defer fmt.Println(i)
	} // here the printing is reversed due to LIFO -> 5 4 3 2 1

	fmt.Println("done..")
}