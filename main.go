package main

import (
	"fmt"

	"github.com/siraken/learn-golang/tog"
)

func main() {
	// A Tour of Go

	callFunctions("Welcome", []func(){
		tog.Welcome1,
		tog.Welcome4,
	})

	callFunctions("Basics", []func(){
		tog.Basics1,
		tog.Basics2,
		tog.Basics3,
		tog.Basics4,
		tog.Basics5,
		tog.Basics6,
		tog.Basics7,
		tog.Basics8,
		tog.Basics9,
		tog.Basics10,
		tog.Basics11,
		tog.Basics12,
		tog.Basics13,
		tog.Basics14,
		tog.Basics15,
		tog.Basics16,
	})

	callFunctions("FlowControl", []func(){
		tog.FlowControl1,
		tog.FlowControl2,
		tog.FlowControl3,
		tog.FlowControl4,
		tog.FlowControl5,
	})
}

func callFunctions(title string, f []func()) {
	for i := 0; i < len(f); i++ {
		fmt.Printf("\n---------- %s %d ----------\n\n", title, i+1)
		f[i]()
	}
}
