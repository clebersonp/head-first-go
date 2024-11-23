package main

import "github.com/clebersonp/head-first-go/ch11/gadget"

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

func main() {
	player := gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playListWithTapPlayers(player, mixTape)
	recorder := gadget.TapeRecorder{}
	playListWithTapRecorders(recorder, mixTape)
}
