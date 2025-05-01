package main

import "fmt"

// Empty interface
interface{}

// In the following example, a "shape" must be able to return its area and perimeter. 
// Both rect and circle fulfill the interface.
type shape interface {
	area() float64
	perimeter() float64
}
  
type rect struct {
	width, height float64
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}
  
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeData(s shape) {
	fmt.Printf("Area: %v - Perimeter: %v\n", s.area(), s.perimeter())
} // prints the area and the perimeter, doesn't matter if it's a circle or a rectangle


// we want to check if s is a circle in order to cast it into its underlying concrete type
// we know s is an instance of the shape interface, but we do not know if it's also a circle
// c is a new circle struct cast from s
// ok is true if s is indeed a circle, or false if s is NOT a circle
c, ok := s.(circle)
if !ok {
	// log an error if s isn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius

// Type Switches:
// A type switch is similar to a regular switch statement, but the cases specify types instead of values.
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}// fmt.Printf("%T\n", v) prints the type of a variable.

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
