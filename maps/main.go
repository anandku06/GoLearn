package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex // assigns nil to the map
// initialied a map using the "map" keyword
// 'm' is the map name : "map" keyword ; datatype of the key in square braces ; datatype of the value to be stored

// using map literals for assigning : assigning the key:value pairs with the assigning
// giving value in the form of key:value pairs
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68438, -74.39967,
	},
	"Google": Vertex{
		37.42282, -122.08408,
	},
}

func main() {
	// initialising using make()
	// m := make(map[string]Vertex) // again initialising the map using the make keyword

	// m["Bell Labs"] = Vertex{
	// 	48.2344535, -45.4563353,
	// } // assigning a key-value pair to the map
	// key acts as index here and value assigned to it

	fmt.Println(m["Bell Labs"])
	fmt.Println(m) // prints the whole map in key-value pairs

	// manipulation of key-value pairs in map 
	m1 := make(map[string]int)

	m1["Answer"] = 42 // initialising a value to a key
	fmt.Println("The value: ", m1["Answer"])
	
	m1["Answer"] = 48 // modifying the value
	fmt.Println("The value: ", m1["Answer"])
	
	delete(m1, "Answer") // deleting the key and value from the map using the delete keyword
	fmt.Println("The value: ", m1["Answer"])

	v, ok := m1["Answer"] // this returns two values, first the value and second the boolean value whether the value is present or not
	fmt.Println("The value: ", v, "Present?", ok)

	for key, value := range m{
		fmt.Println(key, value)
	}
}
