// Arrays

package main

import "fmt"

func main() {

	EvenNum := [5]int{0, 2, 4, 6, 8} // way-1 of initializing array

	/* way-2 of initializing array
		var EvenNum [5]int
		EvenNum[0] = 0
		EvenNum[1] = 2
		EvenNum[2] = 4
		EvenNum[3] = 6
		EvenNum[4] = 8

	fmt.Println(EvenNum[2])
	*/

	for _, value := range EvenNum { // _ means we are not using iterator and value itself is range of numbers
		fmt.Println(value)
	}

	//slicing in arrays

	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[3:5]

	fmt.Println(sliced)

	//define empty slice with data type
	slice2 = make([]int, 5, 10)

	// copying contents of slice

	copy(slice2, numSlice)

	fmt.Println(slice2)

	//appending values to slice

	slice3 := append(numSlice, 3, 0, -1) // adds 3,0,-1 to slice

	fmt.Println(slice3)

}
