package main

import "fmt"

// Go offers a built-in 'recover' function that can stop a program from panicking.
// We'll need to use it to exit the program gracefully.
// There's no point calling recover in the same function as panic, because the panic will continue anyway.
// You can place a call to recover in a separate function, and use defer to call that function before the code that panics.

func calmDown() {
	p := recover()
	// we can try to convert the interface to the type we want via type assertion and ok idiom.
	if err, ok := p.(error); ok {
		fmt.Printf("Recover called. %[1]T: %[1]v\n", err)
	}
	fmt.Println("After recover. This will run.")
}
func freakOut() {
	// This function that panicked will return immediately, and none of the code in that function's block following
	// the panic will be executed.
	defer calmDown() // the deferred function call will recover!
	// We can pass anything to panic because it accepts an empty interface value.
	err := fmt.Errorf("i'm panicking!, there's an error")
	panic(err)
	fmt.Println("After the panic. This won't run.")
}

func main() {
	// When you call 'recover' during normal program execution, it just returns nil and does nothing else.
	fmt.Println("It'll return nil because there's no panic yet:", recover())

	freakOut()
	fmt.Println("Exiting normally.")
}
