package main

import (
	"fmt"

	"github.com/siraken/learn-golang/greet"
)

func main() {

	var name string = "Kento" // this statement can be omitted like this -> name := "Kento"
	fmt.Println(greet.English(name))

	fmt.Printf("%s\n", "Welcome to Golang") // %s for string
	fmt.Printf("%d\n", 10)                  // %d for number

}
