package gadget

import (
	"fmt"
)

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

// Both TapePlayer and TapeRecorder satisfy the Player interface.
// Both can Play a song and Stop a song.

// When a method of defined type has a pointer receiver, we need assign a pointer to the variable of interface type.
// If a type declares methods with pointer receivers, then you'll only be able to use pointers to that type
// when assigning to interface variables.

type Player interface {
	Play(song string)
	Stop()
}

// TryOut is a function that accepts a device of type Player (interface).
// It will call the interface methods on that device.
// And if the device is a TapeRecorder, it will convert it to a TapeRecorder and also call the Record method.
// That conversion is done via reflection.
func TryOut(player Player) {
	fmt.Printf("Trying out Player Type: %T\n", player)
	player.Play("Test Track")
	player.Stop()
	// Type conversions aren't meant for use with interfaces type. So, recorder := TapeRecorder(player) is not allowed
	// We need a 'type assertion' to get the concrete type back.
	// Syntax: var concreteType ConcreteType = interfaceType.(ConcreteType)
	// If the interface type can't be converted to the concrete type, we get a panic at runtime not when compiling.
	// This syntax is also called a 'type assertion' and it has a second return value which is a boolean.
	// The boolean value indicates whether the type assertion succeeded or not and prevents a panic.
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

type TapePlayer struct {
	Batteries string
}

func (t *TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t *TapePlayer) Stop() {
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
