package tog

import (
	"fmt"
	"math"
	"runtime"
	"time"
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
func fc5Sqrt(x float64) string {
	if x < 0 {
		return fc5Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func FlowControl5() {
	fmt.Println(fc5Sqrt(2), fc5Sqrt(-4))
}

// 6. If with a short statement
func FlowControl6() {
	// pow := func(x, n, lim float64) float64 {
	// 	if v := math.Pow(x, n); v < lim {
	// 		return v
	// 	}
	//
	// 	// v cannot be used here
	// 	// return v
	// }

	fmt.Println(
	// pow(3, 2, 10),
	// pow(3, 3, 20),
	)
}

// 7. If and else
func FlowControl7() {
	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %f\n", v, lim)
		}
		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// 8. Exercise: Loops and Functions
func FlowControl8() {
	// sqrt := func(x float64) float64 {
	// }
	//
	// fmt.Println(sqrt(2))
}

// 9. Switch
// break given automatically
func FlowControl9() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// 10. Switch evaluation order
func FlowControl10() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 11. Switch with no condition
func FlowControl11() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() > 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// 12. Defer
func FlowControl12() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// 13. Stacking defers
// LIFO (last-in-first-out): like deck
func FlowControl13() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
