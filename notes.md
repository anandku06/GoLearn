# GOLang Theory and Practice

## Why GO?

- fast, lightweight, and much developer-friendly
- in terms of execution speed, GO is much faster same goes for the compilation speed as well

## Basics

- a go program can be of two types :
  1. Executables -> contains code that can be run ; a program
  2. library/packages -> contains programs that get imported from these libraries or packages, these are built by some one and can be used by other programmers for faster productivity ; must be in _lowercase_ and as short as possible
- all files of GO has an extension of .go (e.g. main.go)

```go
    // basic GO structure

    package main

    import "fmt"

    func main(){
        fmt.Println("Hello World!!");
    }
```

- every GO file must have packaga declaration, by convention each executable GO file must have _main_ package and function
- GO supports UTF-8 character encoding.
- Provides Cross Platform Single Binary executables -> GO can generate executable binaries
- a statically typed language : all datatypes are checked during the program compilation
- **GO's Encapsulation** -> when the functions or variables are accessed declared using capital letter, they can be accessed ouside of thier native package, but not when they are declared with lowercase

```bash
    go run <file_name> # runs the executable file ; more faster than build command
    go build <file_name> # used to build the GO file that returns a executable ; here OS is not specified so the executable is for the current OS

    GOOS=windows GOARCH=amd64 go build . # creates the executable for windows OS with amd54 architecture

    ./main # then running the executable will return the same file
```

## go.mod file

- contains some necessary info regarding the project like the current version of GO and the project's folder name, etc.
- also contains dependencies that the project requires for execution
- generated by using

```bash
# go mod init <projectName>
go mod init github.com/anandku06/GoLearns # in my case, generates the mod file with the package name and the version of GO

go get <package_name> # appends any external package that can be used in the mod file

go mod download # all the packages used in the project is fetched in the mod file

```

## functions

```go
    // syntax for a function in GO
    func functionName(parameter1 datatype, parameter2 datatype, ...) return type{
        // body of the function
    }

    // starts with the keyword "func" and the function name
    // parameters are given like "parameter datatype"
    // after parenthesis the function's return type

    // example
    func add(x int, y int) int {
        return a + b
    }

```

- in GO, parameter datatypes are checked during compilation, so defined datatype must be given as args to avoid errors
- can return multiple values at once

## Variables

- in GO, variables are also declared using the **var** keyword along with the variableName and the datatype.
- in GO, a variable is not uninitialised ever, it gets the default value whenever its declared
- many ways to declare variable in GO, outside and inside the function

```go
    // Syntax -> var varName datatype

    // these types are allowed only outside the main function and are initialised with default values
    var isLoggedIn bool // default value is false
    var x, y, z int // if all the vars are of same type , default value is 0

    func main(){
        var i int // allowed but initialised with its default value
        i := 10 // only allowed inside the function
        var i = 10 // this is also allowed
        var i string = "Hello" // this is the extended version where you fixed the datatype, only string datatype is entertained
    }
```

## Datatypes

- in GO, we have:

  1. Basic -> numbers{int, uint, byte, rune, float, complex}, string, bool
  2. Aggregated -> array, struct
  3. Reference (they don't have the data, they refer the data that are in other places in the program) -> pointers, slices, function, channel, maps
  4. interfaces

- **type conversions** is basically done in statically-typed language like GO, no implicit conversion only explicit conversion
- **type inference** : when compiler infers the type of the data implicitly and this type can't be changed in later phases of the program
- **constants** : means can't change its value once assigned, evaluated during compile time, declared using the _const_ keyword, can store only **basic datatypes**

## Loops

- In GO, we have only for loop, changing the semantics of the for loop syntax, we implement all the other loops
- loops are used to execute some set of instructions again and again until a specified count or condition is met.
- termination condition is important, without it loop is infinte ( can use **break** statement, to stop the execution )

## Conditionals

- Conditionals are like if true, then do this else do this
- use if-else statements for that

## Switch-Case statements

- switch-case statements in GO, just like any other languages
- checks the value of the variable against the cases provided and if the case matches, the code inside that case is executed and the execution exits the switch block
- no fall-through condition, after case matching, it automatically exits the switch block

## defer

- a keyword in GO used to execute a piece of code asynchronously, i.e. parallely to the main execution
- the code with this keyword gets pushed in the exection stack and executed after the main function is executed or returned
- execution follows LIFO
- if any hinderence or any exception occurs in the main function execution, then the defer statements are executed one-by-one
- Usecases can be file handling, database handling, etc. that can hinder the main flow of the program can be exexuted after the main program is returned

## Reference Types (depend on other datatypes)

1. **Pointers** - which points to the address to a variable, contains the address of a variable ; we can access the value on the address by dereferencing it using the asterisk sign(\*)
   - Address here is memory address, virtual address(bcz due to GO runtime, we're not in the actual OS) that contains the value of that variable, when declared automatically assigned an address.
   - accessed using the ampersend sign(&)

## structs

- a type of aggregate datatypes; a collection of different datatypes
- can be accessed using the struct name and the curly braces containing the value for that struct
- fields or data of the struct can be accessed using the dot notation
- default value is default values of its field

```go
// syntax
// type keyword along with the struct name and struct keyword then under the curly braces, we define the data to be provided in the struct
type structName struct{
    // fields / values for this struct
}

// example of a struct that contains the dimensions of a parallelogram
type Sides struct{
    l int // data with its datatype
    b int
}

// to access we call the struct by its name and using the curly braces
// under the braces we provide the values of the same datatype and in the same order they were defined

Sides{24, 45}
```

- fields name under structs must start with capital letters for better exposure of variables outside thier packages
- **structs with pointers** : pointers can also be used with structs, same operations can be performed using pointers as well

- _struct tags_ -> used to give some additional cmds/mapping for values to the struct's fields

```go
    type GenerateCheckoutLink struct{
        PlanID string `json:"planId"` // this is a struct tag
        // here this tag is responsible for parsing the incoming json data's specific key "planId" to the struct field PlanID
    }
```

## arrays in Go

- aggregate datatype used to store multiple datas of same type
- size also is the type of the array
- default value is all the default values of the elements
- due to this the size of an array is fixed because GO is statically-typed, that's the reason it isn't used much

```go
    // syntax
    // var keyword arrayName [size]datatype
    var a [2]string // this means array "a" which is of string and size is of 2 data
    // size of the array is part of its type
    // follows 0-based indexing
```

## slices in GO

- a very special data structure in GO, flexible and useful, also called reference types bcz it has always some backing array from which its values depends
- default value is _nil_
- slicing from the array to create a new array
- main concept is slice doesn't create a new array, it just points to a element that eventually creates a subarray from the original array
- has 3 parts

  1. length -> length of the slice
  2. capacity -> from starting element of the slice to the last element of the backing array
  3. pointer -> to the element of the array through which the slice is made from i.e. the first element of the slice

- slices can also be created without using any backing array
- while its a dynamic array, we don't specify the size, we just keep it empty
- slices can also be made using other slices as well
- we can extend the length of the slice upto the capacity of it
- slice having no elements or not initialised, has 0 length and 0 capacity and is a empty slice (or nil)
- reference types can also be declared using the built-in function **make()**
- **multidimensional slices** : each element of slice is a slice itself

- **append()** : append built-in function _appends elements to the end of a slice_.
  - If it has sufficient capacity, the destination is resliced to accommodate the new elements.
  - If it does not, a new underlying array will be allocated. Append returns the updated slice.
  - It is therefore necessary to store the result of append, often in the variable holding the slice itself,
  - if the capacity isn't enough then, for the future appends it just doubles its capacity
  - Its a _variadic function_ i.e. can take any number of arguments
  - It can be prevented by using the **make()** function

```go
    nums := [5]int {1, 2, 3, 4, 5}

    var half []int = nums[0:3] // here the nums array is sliced from 0th to 2nd(3 not included) position -> 1 2,3 is returned
    // slice is done using the [] braces and under it we give the starting index and the ending index which is exclusive
```

## Maps

- a reference type of default value of **nil**
- GO offers the implementation of hashmaps, dictionaries, etc. in key-value pairs ; key must be comparable ; keys are mostly of strings
- must used to enhance searching

```go
    var m map[string]Vertex
    // package level so using var
    // m is the map name
    // map keyword for map initialisation
    // datatype of keys inside square braces
    // datatype of value to be stored

    m["key"] = value // assigning key-value pair

    // using map literals for initialising
    // assigning the key : value pairs in this form
    var m = map[string]Vertex{
        keys : values ...
    }
```

- intialisation of maps using var is avoided as it assigns _nil_ value to the map and the key-value can't be assigned again to it

### Maps vs Structs

- cannot range loop on structs but possible in maps
- structs fields has thier own memory location but not in the case of maps
- order of the printing of map elements are randomised when looped over
- structs for storing and maps for lookups and faster searching

## Anonymous Functions

- functions without name, or functions declared without giving them a function name ; not follows the particular function declaration
- since functions are considered as first-class citizens, we can assign functions to any variable
- can also be used as parameters in functions

```go
    // Traditional function declaration
    func add(x, y int) int{
        return x + y
    }

    // anonymous function declaration
    hypot := func(x, y float64) float64{
        return math.Sqrt(x * x + y * y)
    }

    // here we used the func keyword to declare the function but not given the function name
    // so the variable containing the function is the function name

    hypot(3, 4) // this is how the function will be called
```

## Closures in GO

- same working as JS
- in case of nested functions, the variable state of outer function is preserved by the inner function
- basically used in place of static keyword in Non-OOPS language

```go
    // closures in GO
    func adder() func(int) int {
        sum := 0
        return func(x int) int {
            sum += x
            return sum
        }
    }
```

## type Keyword
- The type keyword in Go is used to define custom types, type aliases, and struct types. It enhances code readability and maintainability by giving meaningful names to types.
```go
type Age int // defining custom type

type Number = int // defining type aliase

```

## Methods in GO

- A method in Go is a function with a special parameter called a receiver. The receiver ties the function to a specific type, allowing you to call the method on instances of that type using dot notation
- feature in GO that allows us to assign types to functions
- like OOP concept, methods of a class that can access all the properties of its parent class
- used to deifne a certain behaviour or properties for a particular type
- can be used on any datatype
```go
type ReceiverType struct{
    // params for this type struct
}

func (receiver ReceiverType) MethodName(params) returnType {
    // Method body
}

// receiver is the name assigned to the ReceiverType
```
- two types of Receiving type in GO:
    1. Value receiving: 
        - the method uses the copies of the receiver type
        - changes made to the receiver inside the method don't affect the original value
    2. Pointer receiving:
        - method gets the a pointer to the original value
        - changed made to the receiver inside the method affect the original value
- when changes has to be done on a vast number of types, use pointer receiving
- when mutating the values use pointer receiving type
- convention is to make either all the types pointer receiving or value receiving, not both 

## Interfaces in GO
- a type that defines a set of method signatures without specifying how these methods are implemented
- defined using `interface` keyword and contains methods signatures inside it
```go
type Writer interface{
    Write(data string) error
}
// Writer interface having a method that takes s string and returns an error
```
- defines a behaviour a type should have to implement the interface
- methods in GO implicitly satisfies the conditions of interfaces, if not they are not implemented
- default value of an interface is **nil**

- **empty interface** : a interface with no methods, since it has no methods signatures, so it satisfies all the types in GO

- **Type Assertion** : 
    - a mechanism to extract or verify the underlying concrete value of an interface variable.
    - especially for empty interface(interface{}) that can hold any type of data.
    - it lets you check or convert that value to a specific type.
    - Syntax => `value := interface.(datatype)` this is Single-value form
        - returns the value if satisfies the datatype, and stores in the value variable
    - Syntax => `value, ok := interface(datatype)` this is Double-value form, (mostly used)
        - returns two values, if datatype satisfies then, value and true boolean value, else false and that datatypes default value
    - to avoid this writing repeatedly we use **type-switch**
```go
var i interface{} = "hello"

s := i.(string) // returns value, and true as satisfies

f, ok := i.(float64) // returns false and zero as flaot64 can't be assigned to string

f = i.(float64) // creates panic

// Type-Switch : using Switch-case logic to assert the type of the interface
switch value := interface.(datatype){
    case datatype:
        // do somthing
    
    default:
        // do something
}
```