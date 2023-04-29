// Variables and constants

package main

import "fmt"

func main() {
	// A constant is. a variable with a value that can't be changed
	const pi float64 = 3.14

	fmt.Println(pi)

	// declaring multiple variables
	var (
		varA = 2
		varB = 3
	)

	x,y := 14,15 // this will create var x with value 14 and y with value 15 

	fmt.Println(varA, varB)

	// strings are a series of characters between double quotes or backtikcs

	var Name string = "Harshit Singh"

	// get length of string
	fmt.Println(len(Name))

}
