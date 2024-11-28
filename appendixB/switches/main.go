package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize(number int) (string, error) {
	// Go automatically exits the switch at the end of a case's code
	// We don't need to break each case.
	// There's a 'fallthrough' keyword that can be used if you want the next case's code to run as well.
	switch number {
	case 1:
		return "You win a cruise!", nil
	case 2:
		return "You win a car!", nil
	case 3:
		return "You win a goat!", nil
	default:
		return "", fmt.Errorf("invalid door number %d", number)
	}
}

func main() {
	rand.NewSource(time.Now().Unix())
	randomNumber := rand.Intn(3) + 1
	phrase, err := awardPrize(randomNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(phrase)
}
