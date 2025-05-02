package main

import (
	"fmt"
	"errors"
)

// Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. 
// Maps are a data structure that provides key->value mapping.
// The zero value of a map is nil.
// We can create a map by using a literal or by using the make() function:
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

// The len() function works on a map, it returns the total number of key/value pairs.
ages = map[string]int{
	"John": 37,
	"Mary": 21,
}
fmt.Println(len(ages)) // 2

// Map of key = string, value = struct
users := make(map[string]user)
users["John"] = user{
	name:        "John",
	phoneNumber: 3335564812,
}

type user struct {
	name        string
	phoneNumber int
}

// Mutations:
// Insert an Element
m[key] = elem
// Get an Element
elem = m[key]
// Delete an Element
delete(m, key)
// Check If a Key Exists
elem, ok := m[key]
// If key is in m, then ok is true and elem is the value as expected.
// If key is not in the map, then ok is false and elem is the zero value for the map's element type.
