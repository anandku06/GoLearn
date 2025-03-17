package main

import (
	"fmt"
	"math"
)

// methods in GO
// defining a receiving type that a function will use (struct)
type Vertex struct{
	x, y float64
}

type MyFloat float64

// giving the instance of the receiving type to this function nameed v
// different from the regular function by assigning the receiving type before the function name
func (v Vertex) Abs() float64{
	return math.Sqrt(v.x * v.x + v.y * v.y) // here accessing the field of the struct type and performing the operation
}
// this is also called Value Receiving : method gets the copy of the original value here

func Abs(v Vertex) float64{ // this is a function not a method
	return math.Sqrt(v.x * v.x + v.y * v.y) // here accessing the field of the struct type and performing the operation
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f) // bcz MyFloat is now a new type that can't be implicitly be type casted
		// so explicitly type casting it to match the return type
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64){
	v.x = v.x * f
	v.y = v.y * f
}
// this is Pointer Receiving : method gets the pointer to the original value i.e. if change happens then it'll affect the original value as well

func main() {
	v := Vertex{x: 3, y: 4} // declaring a var of struct
	fmt.Println(v.Abs()) // we are able to use this bcz we given the same datatype to this function and this datatype acts as the receiving type of that function
	// here v is given to that function as its receiving type and also the values given to the function

	fmt.Println(Abs(v)) // this is a regular function call that takes a struct as an argument

	f := MyFloat(-math.Sqrt(4))
	fmt.Println(f.Abs())
	v.Scale(10) // now the value is changed here
	fmt.Println(v.x, v.y) // here the value is changed
	fmt.Println(v.Abs()) // after changing the value
}