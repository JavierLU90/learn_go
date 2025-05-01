package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// One way to declare variables:
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// Second way to declare variables (more dynamic):
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// Declare Constants:
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	// Formatting Strings:
	s := fmt.Sprintf("I am %v years old", 10) // Default
	s := fmt.Sprintf("I am %s years old", "way too many") // String
	s := fmt.Sprintf("I am %d years old", 10) // Int
	s := fmt.Sprintf("I am %f years old", 10.523) // Float
	s := fmt.Sprintf("I am %.2f years old", 10.523) // Float (Rounded)

	fmt.Print(s)
}
