package main

// when multiple imports they are enclosed inside parenthesis
import (
	"fmt"
	"math"
	"math/rand" // here importing rand package from math package
)

func main(){
	fmt.Println("My favourite number is -> ", rand.Intn(10)) // genrates a random number from 0 to 10
	fmt.Println(math.Pi) // returns the PI value ; accessing a variable from the package math
}