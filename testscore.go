package main

func ex_check_score() {
	checkScore(85)
	checkScore(50)
}

func checkScore(score int) {
	isPassed := 80 < score

	if isPassed {
		println("Passed:", isPassed)
	} else {
		println("Failed:", isPassed)
	}
}
