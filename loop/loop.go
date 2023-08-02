package loop

import "fmt"

func CountNumbers(first, last int) string {
	var str string

	for i := first; i <= last; i++ {
		if i == last {
			str += fmt.Sprintf("%d", i)
		} else {
			str += fmt.Sprintf("%d, ", i)
		}
	}

	return str
}
