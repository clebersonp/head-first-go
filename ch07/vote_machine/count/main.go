package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch07/vote_machine_v1/datafile"
	"log"
)

// To run on terminal:
// go run ./ch07/vote_machine_v1/count/main.go

func main() {
	lines, err := datafile.GetStrings("ch07/vote_machine_v1/count/votes.txt")
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
		votes[line]++
	}
	fmt.Println("Count votes with maps:")
	for candidate, count := range votes {
		fmt.Printf("%s: %d\n", candidate, count)
	}
}
