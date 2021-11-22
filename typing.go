package main

import (
	"fmt"
)

func main() {

	// total
	var totalScore int = 0

	// questions
	ask(1, "hello", &totalScore)
	ask(2, "bonjour", &totalScore)
	ask(3, "こんにちは", &totalScore)

	// show result
	fmt.Printf("Result: %d\n", totalScore)
}

func ask(quizNo int, content string, scorePtr *int) {

	// user input
	var input string

	// show question and get user input
	fmt.Printf("Question %d: %s\n", quizNo, content)
	fmt.Scan(&input)

	// score
	if input == content {
		fmt.Println("Correct!")
		*scorePtr += 10
	} else {
		fmt.Println("Incorrect :(")
	}
	
}
