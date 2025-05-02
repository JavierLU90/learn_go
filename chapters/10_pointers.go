package main

// Pointers:
// A pointer is a variable that stores the memory address of another variable. 
// This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.
// The * syntax defines a pointer:
var p *int

//The & operator generates a pointer to its operand.
myString := "hello" // myString is just a string
myStringPtr := &myString // myStringPtr is a pointer to myString's address

// Dereference:
*myStringPtr = "world"                              // set myString through the pointer
fmt.Printf("value of myString: %s\n", *myStringPtr) // read myString through the pointer
// value of myString: world

// Example:
// One of the most common use cases for pointers in Go is to pass variables by reference, 
// meaning that the function receives the address of the original variable, not a copy of the value. 
// This allows the function to update the original variable's value.
func increment(x *int) {
    *x++
    fmt.Println(*x) // 6
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x) // 6
}

// Empty pointer (nil pointers):
var p *int
fmt.Printf("value of p: %v\n", p)// value of p: <nil>
// If a pointer points to nothing (the zero value of the pointer type) 
// then dereferencing it will cause a runtime error (a panic) that crashes the program. 
// Generally speaking, whenever you're dealing with pointers you should check if it's nil 
// before trying to dereference it.

// Pointer Receiver Code:
// Methods with pointer receivers don't require that a pointer is used to call the method. 
// The pointer will automatically be derived from the value.
type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow() {
    c.radius *= 2
}

func main() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius) // prints 8
}
