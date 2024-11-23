package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch12/sum"
	"github.com/clebersonp/head-first-go/ch12/sum/file"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	numbers, err := file.GetFloats(fileName)
	if err != nil {
		log.Fatal(err)
	}
	total := sum.Plus(numbers)
	fmt.Printf("Total: %0.2f\n", total)
}
