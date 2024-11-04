package main

func main() {

	// Avoid shadowing names wherever possible. Choosing non-conflicting names for the variables.

	// The bellow examples is bad practice of shadowing and do not compile at all.

	// int name shadowing the int type and will take precedence over it
	var int int = 12
	// append name shadowing the function append and will take precedence over it
	var append string = "minutes of bonus footage"
	// fmt name shadowing the fmt package and will take precedence over it
	var fmt string = "DVD"

	var count int // int it's not int type anymore. Now the int name will be a simple variable name, not type anymore
	var languages = append([]string{}, "English")
	fmt.Println(int, append, "on", fmt, languages, count)

}
