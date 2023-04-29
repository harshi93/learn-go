// Defer, Recover, Panic and multiargs function

package main

import "fmt"

func main() {

	/*
	   Panic is similar to exceptions in other language
	   Recover is used regain normal flow after a panic,
	   Defer functions are run as a result of recover
	   Defer is used as a cleanup mechanism or wait mechanism
	   unit all surrounding functions are executed and anything
	   deferred function gets pushed to last of stack
	*/

	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
	demPanic()

	fmt.Println(multiargsfuncdemo(10, 20, 30, 40, 50, 90))

	//defer FirstRun()
	//SecondRun()

}

func multiargsfuncdemo(args ...int) int {
	sum := 0

	for _, value := range args {
		sum += value
	}

	return sum
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2

	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}

/*
func FirstRun() {
	fmt.Println("I am very excited")
}

func SecondRun() {
	fmt.Println("I am not so excited")
}
*/
