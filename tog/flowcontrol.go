package tog

import (
	"fmt"
)

// 1. For
func FlowControl1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 2. For continued
// * The init and post statements are optional.
func FlowControl2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 3. For is Go's "while"
func FlowControl3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 4. Forever
func FlowControl4() {
	// This loop will spin, using full CPU
	fmt.Println("for { }")
}

// 5. If
func FlowControl5() {
	// sqrt := func(x float64) string {
	// 	if x < 0 {
	// 		return sqrt(-x) + "i"
	// 	}
	// 	return fmt.Sprint(math.Sqrt(x))
	// }
	// fmt.Println(sqrt(2), sqrt(-4))
}
