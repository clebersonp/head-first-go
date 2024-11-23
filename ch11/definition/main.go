package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch11/gadget"
)

// playList is a function that accepts a device of type TapePlayer and a slice of strings.
// We are passing the concrete type gadget.TapePlayer as the first parameter, and the slice of songs as the second parameter.
func playListWithTapPlayers(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func playListWithTapRecorders(device gadget.TapeRecorder, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// playList is a function that accepts a device of type Player (interface) and a slice of strings.
// Any type that has all the methods listed in an interface definition is said to 'satisfy' that interface
// can be passed here.
func playList(device gadget.Player, songs []string) {
	fmt.Printf("Playing with interface: %T\n", device)
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	tapePlayer := gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playListWithTapPlayers(tapePlayer, mixTape)
	tapeRecorder := gadget.TapeRecorder{}
	playListWithTapRecorders(tapeRecorder, mixTape)

	fmt.Println()
	// we can pass the concrete type
	// to satisfy the interface we need pass the pointer to the concrete type,
	// because the concrete type has all the methods defined in the interface with pointer receiver
	playList(&tapePlayer, mixTape)
	playList(tapeRecorder, mixTape)

	fmt.Println()
	// we can use a variable of interface type and assign to it the concrete type
	var player gadget.Player
	player = &gadget.TapePlayer{}
	playList(player, mixTape)
	player = gadget.TapeRecorder{}
	playList(player, mixTape)
}
