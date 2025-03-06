// switch-case statements in GO
package main

import (
	"fmt"
	"runtime" // contains functions like GOOS, GOARCH, etc.
	"time"    // contains functions related to time
)

func main()  {
	fmt.Println("Go runs on")

	// here as well initialisation with condition checking
	// under the switch statement, os variable is initialised and this variable's scope is inside the switch only
	switch os := runtime.GOOS; os { // checks the value of the OS
	case "darwin": // if the case matches, then this part is executed and execution breaks out of this scope
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default: // if none of the case matches, then this will execute
		fmt.Printf("%s.\n", os)
	}

	fmt.Println(runtime.GOOS, runtime.GOARCH) // returns the current OS and ARCHITECTURE of the OS ; in the runtime package 


	today := time.Now().Weekday() // Now -> returns the current instance of local time
	// Weekday returns the day of the week
	fmt.Println(today) 

	switch time.Saturday{
	case today + 0: // if added any int value it shifts the day to the next coming day according to the current value
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")

	case today + 2:
		fmt.Println("In two days")
	
	default:
		fmt.Println("Too far!!")
	}

	// switch statement without any condtion or variable
	t := time.Now()

	// here this switch is acting like if-else block, not hecking any variable value but case are like if-elseif-else
	switch{
	case t.Hour() < 12:
		fmt.Println("Good Morning!!")

	case t.Hour() < 17:
		fmt.Println("Good Afternoon!!")

	default:
		fmt.Println("Good Evening!!")
	}
}