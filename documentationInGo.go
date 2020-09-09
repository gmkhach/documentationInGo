// In Go documentation is built right into the comments.
// These two lines will print in your terminal if you navigate to the directory where this file is and type-in the "go doc" command.
package main

import (
    "fmt"
)

func main() {
	a, b := 21, 21
	fmt.Println(IntSum(a, b))
}

// Sum takes two integers and reterns an integer (the sum of the two arguments).
// Type-in "go doc IntSum" in the terminal to read this comment, as well as the func signature.
func IntSum(a, b int) int {
	return a + b;
}