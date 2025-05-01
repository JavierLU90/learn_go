package main

import (
	"fmt"
	"errors"
)

// Declaring Arrays
func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{
		primary,
		secondary,
		tertiary,
	}
	lengths := [3]int{
		len(primary),
		len(primary) + len(secondary),
		len(primary) + len(secondary) + len(tertiary),
	}
	return messages, lengths
}

// Slices:
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4] // mySlice = {3, 5, 7}

// Most of the time we don't need to think about the underlying array of a slice. 
// We can create a new slice using the make function:
// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)
// Slices created with make will be filled with the zero value of the type.
// If we want to create a slice with a specific set of values, we can use a slice literal:
mySlice := []string{"I", "love", "go"}
// Make an empty slice:
mySlice := []float64{}


// Variadic Functions:
// Many functions, especially those in the standard library, 
// can take an arbitrary number of final arguments. 
// This is accomplished by using the "..." syntax in the function signature. 
// A variadic function receives the variadic arguments as a slice.
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}

// Spread Operator:
// The spread operator allows us to pass a slice into a variadic function. 
// The spread operator consists of three dots following the slice in the function call.
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

// Append:
// The built-in append function is used to dynamically add elements to a slice:
func append(slice []Type, elems ...Type) []Type
// If the underlying array is not large enough, 
// append() will create a new underlying array and point the returned slice to it.
// Notice that append() is variadic, the following are all valid:
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

// Range:
// for INDEX, ELEMENT := range SLICE {}
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}

// Slice of Slices:
// Slices can hold other slices, effectively creating a matrix, or a 2D slice.
matrix := [][]int{} // create empty matrix of unknown size
matrix := make([][]int, rows) // create empty matrix of known rows
