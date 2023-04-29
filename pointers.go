// Call by reference can't be performed without pointers

package main

import "fmt"

func main() {
	x := 5


	changeValue(&x) // this way of passing value ensures that variable value is changed at memory address
	fmt.Println(x)
	

}

func changeValue(x *int) {
	*x = 1000
}