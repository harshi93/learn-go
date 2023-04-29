// Decision Making / Conditionals

package main

import "fmt"

func main() {

	var age int = 18
	/*

		if age >= 18 {
			fmt.Println("you can vote")
		} else {
			fmt.Println("you can't vote")
		}
	*/

	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 18:
		fmt.Println("Focus on studies")
	case 20:
		fmt.Println("Get a job")
	default:
		fmt.Println("are you alive ?")
	}
}
