package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch07/vote_machine/datafile"
	"log"
)

// To run on terminal:
// go run ./ch07/vote_machine_v1/count/main.go

func main() {
	lines, err := datafile.GetStrings("ch07/vote_machine/count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	countVotesWithSlices(lines)
	fmt.Println()
	countVotesWithMaps(lines)
}

func countVotesWithSlices(lines []string) {
	// two slices to track the names and their counts of votes
	// It's more complex way to count votes
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
				// I don't need to check the rest of the names
				break
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	fmt.Println("Count votes with two slices:")
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

func countVotesWithMaps(lines []string) {
	// create and initialize a map with built-in make function
	// syntax: map[key-type]value-type
	votes := make(map[string]int)
	for _, line := range lines {
		// Arrays and slices only let you use integer indexes.
		// But you can choose almost any type to use for a map's keys.
		// If the key we're accessing doesn't already exist, we'll get the zero value back (0), then increment it
		// giving us 1. When we access the same key again, we'll get 1 and increment normal.
		votes[line]++
	}
	fmt.Println("Count votes with maps:")
	// syntax: for key, value := range map
	// or: for key := range map
	// or: for _, value := range map
	for candidate, count := range votes {
		fmt.Printf("%s: %d\n", candidate, count)
	}
}
