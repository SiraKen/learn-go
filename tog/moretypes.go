package tog

import "fmt"

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
