package main

import "fmt"

type any interface{}

// Type Assertion 
func main() {
	var i any = "hello" // here this empty interface is given a string

	// asking whether it can hold this string value or not
	s := i.(string) // if possible then returns value, Single-value form
	fmt.Println(s)

	s, ok := i.(string) // if possible, returns value and bool value, Double-value form
	fmt.Println(s, ok)

	f, ok := i.(float64) // returns false , 0, bcz can't possible and due to type assertion, it doesn't throw any error
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)

	do(21)
	do("Hello")
	do(true)
}

// to avoid this repeatedly writing type assertion, use type-switch
func do(i any){
	switch v := i.(type) {
	case int:
		fmt.Printf("This is %T\n", v)
	
	case string:
		fmt.Printf("This is %T\n", v)
	
	default:
		fmt.Println("I don't know about this!!!")
	}
}