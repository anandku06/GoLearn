package main

import "fmt"

// arrays in GO
func main(){
	var a [2]string // declared an array of type [2]string
	a[0] = "Hello" // same how to access and modify the elements of an array using indexing
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}