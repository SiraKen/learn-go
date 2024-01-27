package tog

import (
	"fmt"
	"strings"
)

// 1. Pointers
func Moretypes1() {
	i, j := 42, 2701

	p := &i         // p -> address of i
	fmt.Println(*p) // *p -> value of i

	*p = 21 // 42 -> 21
	fmt.Println(i)

	p = &j       // set address of j (2701) to p
	*p = *p / 37 // j = 2701 / 37

	fmt.Println(j)
}

// 2. Structs
func Moretypes2() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
}

// 3. Struct Fields
func Moretypes3() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)
}

// 4. Pointers to structs
func Moretypes4() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v
	// I can write as below instead of (*p).X = 1e9
	p.X = 1e9

	fmt.Println(v)
}

// 5. Struct Literals
func Moretypes5() {
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)
}

// 6. Arrays
func Moretypes6() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// 7. Slices
func Moretypes7() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 3, 5, 7
	fmt.Println(s)
}

// 8. Slices are like references to arrays
// *** Difficult for me ***
func Moretypes8() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // John, Paul
	b := names[1:3] // Paul, George
	fmt.Println(a, b)

	b[0] = "XXX" // Paul -> XXX
	fmt.Println(a, b)
	fmt.Println(names)
}

// 9. Slice leterals
func Moretypes9() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// 10. Slice defaults
func Moretypes10() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // 3, 5, 7
	fmt.Println(s)

	s = s[:2] // 3, 5
	fmt.Println(s)

	s = s[1:] // 5
	fmt.Println(s)
}

// 11. Slice length and capacity
func Moretypes11() {
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

// 12. Nil slices
// The zero value of a slice is `nil`
func Moretypes12() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// 13. Creating a slice with `make`
func Moretypes13() {
	printSlice := func(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}

	a := make([]int, 5) // len=5 cap=5 [0 0 0 0 0]
	printSlice("a", a)

	b := make([]int, 0, 5) // len=0 cap=5 []
	printSlice("b", b)

	c := b[:2] // len=2 cap=5 [0 0]
	printSlice("c", c)

	d := c[2:5] // len=3 cap=3 [0 0 0]
	printSlice("d", d)
}

// 14. Slices of slices
func Moretypes14() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
