// File I/O

package main

import (
	"fmt"
	"log"
	"os"
)

func fileIO() {

	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This file was created using Golang")

	file.Close()

	stream, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)
}
