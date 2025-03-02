// every GO file must start with the package declaration
// tells that this file belongs to the specified package
package main

import "fmt" // importing the "fmt" package from the standard library
// this package contains functions regarding printing and string manipulation

// this is a function -> a group of logically related code that does a single task ; declared by using "func" keyword

func main() {
	fmt.Println("Hello World!!")
}
// main function is the starting point of the compiler