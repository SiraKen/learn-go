package tog

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// 1. Packages
func Basics1() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// 2. Imports
func Basics2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// 3. Exported names
func Basics3() {
	fmt.Println(math.Pi)
}

// 4. Functions
func Basics4() {
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

// 5. Functions continued
func Basics5() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

// 6. Multiple results
func Basics6() {
	swap := func(x, y string) (string, string) {
		return y, x
	}

	a, b := swap("hello", "world")

	fmt.Println(a, b)
}

// 7. Named return values
func Basics7() {
	split := func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		// naked return:
		return
	}

	fmt.Println(split(17))
}

// 8. Variables
var b8i int

func Basics8() {
	var c, python, java bool
	fmt.Println(b8i, c, python, java)
}

// 9. Variables with initializers
var b9i, b9j int = 1, 2

func Basics9() {
	c, python, java := true, false, "no!"
	fmt.Println(b9i, b9j, c, python, java)
}

// 10. Short variable declarations
func Basics10() {
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
}

// 11. Basic types
//
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// alias for uint8:
// byte
//
// alias for int32:
// rune
// (represents a Unicode code point)
//
// float32 float64
//
// complex64 complex128
func Basics11() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// 12. Zero values
func Basics12() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 13. Type conversions
func Basics13() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 14. Type inference
func Basics14() {
	// v := 42 -> int
	// v := 42.123 -> float64
	v := "hello"
	fmt.Printf("v is of type %T\n", v)
}

// 15. Constants
// * Constants cannot be declared using the := syntax.
const B15Pi = 3.14

func Basics15() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", B15Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// 16. Numeric Constants

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func Basics16() {
	needInt := func(x int) int { return x*10 + 1 }
	needFloat := func(x float64) float64 {
		return x * 0.1
	}

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
