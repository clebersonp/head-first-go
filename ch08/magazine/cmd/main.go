package main

import (
	"fmt"
	// import path to import other package, magazine in this case
	// All exported names from the magazine package are available. It'll can be struct, function, constants, methods and so on.
	"github.com/clebersonp/head-first-go/ch08/magazine"
)

func main() {
	// We can access exported names from the magazine package because its name is capitalized (exported)
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
