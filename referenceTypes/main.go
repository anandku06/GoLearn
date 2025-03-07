package main

import "fmt"

func main(){
	i, j := 42, 2701

	p := &i // points to i, storing the address of i in p using the & sign
	fmt.Println(*p) // read i through the pointer, dereferencing
	*p = 21 // set i through the pointer, bcz it has the value, it can be modified
	fmt.Println(i) // now i's value is changed, bcz it has the reference to the original value so, the original value is also changed

	p = &j // point to j, same here as well stored address of j to p
	*p /= 37 // changed the value of j here by dividing the pointer
	fmt.Println(j) // now the original value is changed as well
}