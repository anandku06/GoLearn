package main

import (
	"fmt"
	"math"
)

// declared an interface Abser containing a method signature for Abs()
type Abser interface{
	Abs() float64
}

type any interface{} // defining an type alias for empty interface

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser 
	// this f assignment is ok, bcz MyFloat satisfies the method Abs() in the interface, i.e. return type 
	a = &v // a *Vertex implements Abserf 

	// in this line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser
	// a = v // bcz Abs() method is assigned to the pointer receiver not value receiver

	fmt.Println(a.Abs())

	// var i interface{} // this is an empty interface ; anything can be assigned to it
	var i any // using the type alias of interface here
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i)
}

type MyFloat float64

// another method of Value receiving
func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// a struct Vertex
type Vertex struct{
	x, y float64
}

// a method with Pointer receiving
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.x * v.x + v.y * v.y)
}