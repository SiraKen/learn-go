package main

import (
	"fmt"
)

func main() {

	var name string = "Kento" // this statement can be ommited like this -> name := "Kento"
	println("Hello, " + name)

	fmt.Println("Hello, Go")

	fmt.Printf("%s\n", "Welcome to Golang") // %s for string
	fmt.Printf("%d\n", 10) // %d for number

}
