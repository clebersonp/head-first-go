package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch14/prose"
)

func main() {
	fruits := []string{"apple", "banana"}
	fmt.Println("I need to by", prose.JoinWithCommas(fruits))
	fruits = []string{"water melon", "pineapple", "melon"}
	fmt.Println("I need to by", prose.JoinWithCommas(fruits))
}
