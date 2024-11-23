package gadget

import "fmt"

// Interface: Two different types that have the same methods. We can define a contract such as an interface.
// An interface is a set of methods that certain values are expected to have (implement).
// Any type that has all the methods listed in an interface definition is said to 'satisfy' that interface.
// A type that satisfies an interface can be used anywhere that interface is called for.
// A type can have methods in addition to those listed in the interface, but it mustn't be missing any,
// or it doesn't satisfy that interface.
// A type can satisfy multiple interfaces, and an interface can (and usually should) have multiple types that satisfy it.
// Many other languages would require us to explicitly say that SomeType satisfies SomeInterface.
// But in Go, this happens 'automatically'.

// A concrete type specify the underlying type that holds the value's data.
// Interface types don't have underlying types. They only describe what a value can do (what methods it has).
// Syntax: type InterfaceName interface { list of methods }
// When you have a value in a variable with an interface type, you can only call methods defined as part of that interface.

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
