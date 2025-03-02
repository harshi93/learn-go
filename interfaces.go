// Interfaces

package main

import (
	"fmt"
	"math"
)

func mainInterfaces() {

	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle:", getArea(rect))
	fmt.Println("Area of circle:", getArea(circ))
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func mainInterfacesCaller() {
	mainInterfaces()
}
