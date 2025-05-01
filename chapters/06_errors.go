package main

import (
	"fmt"
	"errors"
)

// When something can go wrong in a function, 
// that function should return an error as its last return value. 
// Any code that calls a function that can return an error should handle errors 
// by testing whether the error is nil.

func Atoi(s string) (int, error)

// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then the
// variable i was converted successfully


// Because errors are just interfaces, 
// you can build your own custom types that implement the error interface.
type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
    return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
		// you can also do this with the errors library
		err := errors.New("no dividing by 0")
		return 0, err
	}
	return dividend / divisor, nil
}
