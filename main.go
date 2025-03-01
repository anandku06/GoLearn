// every GO file must start with the package declaration
// tells that this file belongs to the specified package
package main 

import "fmt" // importing the "fmt" package from the standard library
// this package contains functions regarding printing and string manipulation

// this is a function -> a group of logically related code that does a single task ; declared by using "func" keyword
func main(){
	fmt.Println("Hello World!!") // using the package "fmt" and by using the dot(.) we access the functions inside that package
	// here the function name's first letter is captial
	// this is the appraoch used by GO to signify which functions are exposed outside their package
}
// main function is the starting point of the compiler