// Call by reference can't be performed without pointers

package main

import "fmt"

func mainPointers() {
	x := 5

	changeValue(&x) // this way of passing value ensures that variable value is changed at memory address
	fmt.Println(x)
}

func mainPointersCaller() {
	mainPointers()
}

func changeValue(x *int) {
	*x = 1000
}
