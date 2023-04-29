package main

import "fmt"

func main() {

	for i := 1; i <= 15; i++ {

		fmt.Println(i)

	}

	whileLoopLikeForLoop()

}

func whileLoopLikeForLoop() {
	i := 1

	for i <= 10 {
		fmt.Println("This is while loop like for loop", i)
		i++
	}
}
package main

import "fmt"

func main() {

	for i := 1; i <= 15; i++ {

		fmt.Println(i)

	}

	whileLoopLikeForLoop()

	nestedFor()

}

func whileLoopLikeForLoop() {
	i := 1

	for i <= 10 {
		fmt.Println("This is while loop like for loop", i)
		i++
	}
}

func nestedFor() {
	for i := 1; i < 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
