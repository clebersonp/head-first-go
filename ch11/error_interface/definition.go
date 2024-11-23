package main

import "fmt"

// An error value is any value with a method named 'Error' that returns a string.
// That error value is of type 'error', and any concrete error type must implement this method.

// ComedyError is a defined type. Because it has an Error method that returns a string, it satisfies the error interface,
// and we can assign it to a variable with the type error.
type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

// If you need an error value, but also need to track more information about the error than just an error message string,
// you can create your own type that satisfies the error interface and stores the information you want.

// OverheatError uses float64 as its underlying type, allowing us to track the degrees over capacity.
type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees", o)
}

// checkTemperature checks if the actual temperature exceeds the safe temperature limit.
// It takes two float64 parameters:
// - actual: the current temperature reading
// - safe: the maximum safe temperature threshold
// Returns an OverheatError if actual temperature exceeds safe temperature, nil otherwise.
func checkTemperature(actual, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = fmt.Errorf("this is an example of error type")
	fmt.Printf("Error type: %[1]T, value: %[1]v, Error() method: %[2]s\n", err, err.Error())

	err = ComedyError("this is an example of custom error type")
	fmt.Printf("Error type: %[1]T, value: %[1]v, Error() method: %[2]s\n", err, err.Error())

	err = checkTemperature(60.0, 35.0)
	if err != nil {
		fmt.Printf("Error type: %[1]T, value: %[1]v, Error() method: %[2]s\n", err, err.Error())
	}

}
