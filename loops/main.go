package main

import "fmt"

func main() {
	// concept of for loop

	sum := 0                  // declared a variable
	for i := 0; i < 10; i++ { // typical for loop syntax : initialisation, condition, updation
		sum += i
	}
	fmt.Println(sum)

	// while loop using for syntax
	j := 0
	for j < 10 { // just like while, the loop will run until the condition is true
		fmt.Print(j)
		j++

		if j == 5 {
			break // stops the execution of the loop, exits the loop
		}
	}
}

// for initialisation; condition; updation {
// 	code to execute
// }
