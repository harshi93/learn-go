package main

import "fmt"

func arithmeticOperators() {
	const pi float64 = 3.16412732
	x := 5
	isbool := true
	var name string = "Harshit Singh"

	fmt.Printf(" %f \n", pi)     // prints floating point, set the precision by saying .3f instead of f
	fmt.Printf(" %T \n", name)   // prints type of anything
	fmt.Printf(" %t \n", isbool) //prints boolean
	fmt.Printf(" %d \n", x)      // prints integers
	fmt.Printf(" %b \n", 25)     //prints binary
	fmt.Printf(" %c \n", 33)     // print character associate with keycode
	fmt.Printf(" %x \n", 15)     //prints hexcode
	fmt.Printf(" %e \n", pi)     // prints values with scientific notations
}
