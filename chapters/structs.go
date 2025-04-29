package main

// Declaring Structs
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

// Anonymous Structs
// Like a normal struct,
// but it is defined without a name and therefore cannot be referenced elsewhere in the code.
myCar := struct {
	brand string
	model string
} {
	brand: "Toyota",
	model: "Camry",
}

// Embedded Structs:
lanesTruck := truck{
	bedSize: 10,
	car: car{
	  brand: "Toyota",
	  model: "Camry",
	},
}
  
fmt.Println(lanesTruck.brand) // Toyota
fmt.Println(lanesTruck.model) // Camry

// Struct Methods:
type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}
  
var r = rect{
	width: 5,
	height: 10,
}
  
fmt.Println(r.area())// prints 50
