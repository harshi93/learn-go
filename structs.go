// Structures

package main

import "fmt"

func main() {

	rect1 := Rectangle{10, 5}
	fmt.Println(rect1.height)

	fmt.Println("Area of rectangle is", rect1.area())

}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
