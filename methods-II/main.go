package main

import (
	"fmt"
)

// difference between traditional function and methods

type Vertex struct {
	x, y float64
}

// method with pointer receiving type
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// regular function with a pointer parameter
func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	// &v.Scale(3) // this is not required bcz GO implicitly do this thing when called
	v.Scale(2) // here method doesn't require & sign as it requires pointer type

	// ScaleFunc(v, 10) // here, vice-versa : throws error as the parameter requires pointer not the value
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3} // this time using the pointer of the receiving type ; but not needed as GO implicitly does this on its own
	p.Scale(3)
	ScaleFunc(p, 4)
}
