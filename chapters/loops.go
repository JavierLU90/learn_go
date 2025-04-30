package main

import "fmt"

// For Loops:
for i := 0; i < 10; i++ {
	fmt.Println(i) // Prints 0 through 9
}

// Loops in Go can omit sections of a for loop. 
// For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
func maxMessages(thresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost > thresh {
			return i // returns the max number of messages the user can send
		}
	}
}

// Most programming languages have a concept of a while loop. 
// Because Go allows for the omission of sections of a for loop, 
// a while loop is just a for loop that only has a CONDITION.
plantHeight := 1
for plantHeight < 5 {
	fmt.Println("still growing! current height:", plantHeight) // prints the current height
	plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches") // prints the final height

// Continue and break:
// The continue keyword stops the current iteration of a loop and continues to the next iteration. 
// continue is a powerful way to use the guard clause pattern within loops.
for i := 0; i < 10; i++ {
	if i % 2 == 0 {
	  continue
	}
	fmt.Println(i) // only prints odd numbers
}
// The break keyword stops the current iteration of a loop and exits the loop.
for i := 0; i < 10; i++ {
	if i == 5 {
	  break
	}
	fmt.Println(i) // prints from 0 to 4 only
}
