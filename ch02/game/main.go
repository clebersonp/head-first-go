// guess challenges players to guess a random number.
package main

import (
	"fmt"
	"log"
)

func main() {
	start, end := 1, 100
	fmt.Print("Type the number of guesses you want to take: ")
	numberGuesses, err := getUserInput()
	if err != nil {
		log.Fatal(err)
	}
	startGame(start, end, numberGuesses, randomNumber(end))
}

func startGame(start, end, numberGuesses, target int) {
	fmt.Println("Guess a number between", start, "and", end, ".")
	fmt.Println("Type your guesses below.")
	var success bool
	for i := numberGuesses; 0 < i; i-- {
		fmt.Println("You have", i, "guesses left")
		userGuess, err := getUserInput()
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		if userGuess == target {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		} else if userGuess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Oops. Your guess was HIGH.")
		}
	}
	if !success {
		fmt.Println("Sorry. You ran out", numberGuesses, "of guesses and didn't guess my number. It was:", target)
	}
}
