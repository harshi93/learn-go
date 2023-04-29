// Functions and Recursion

package main

import "fmt"

func main() {

	x, y := 5, 6

	fmt.Println(add(x, y))

	num := y

	fmt.Println(factorial(num))

}

func add(num1, num2 int) int {
	// int on the outside is type for value returned
	return num1 + num2
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
