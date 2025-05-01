package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	test("Lane,", " happy birthday!")
	test("Zuck,", " hope that Metaverse thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}


// Ignore return values:
func getPoint() (x int, y int) {
	return 3, 4
}

x, _ := getPoint()// ignore y value


// Go functions can return naked values:
func getCoords() (x, y int){
	// x and y are initialized with zero values
	return // automatically returns x and y
} // Naked returns should only be used in short and simple functions.


/* Defer
Example: 
No matter where it returns, it should print the following message just before it returns:
TEXTIO BOOTUP DONE
Using defer you only have to write this message once.

defer fmt.Println("TEXTIO BOOTUP DONE") */

// Anonymous Functions:
newX, newY, newZ = conversions(func(a int) int {
	return a + a
}, 1, 2, 3)
// newX is 2, newY is 4, newZ is 6
