package main

func checkScore(score int) {
	isPassed := 80 < score

	if isPassed {
		println("Passed:", isPassed)
	} else {
		println("Failed:", isPassed)
	}
}

func main() {
	checkScore(85)
	checkScore(50)
}
