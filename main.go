package main

import (
	"fmt"

	"github.com/siraken/learn-golang/tog"
)

func main() {
	// A Tour of Go

	// callFunctions("Welcome", []func(){
	// 	tog.Welcome1,
	// 	tog.Welcome4,
	// })

	// callFunctions("Basics", []func(){
	// 	tog.Basics1,
	// 	tog.Basics2,
	// 	tog.Basics3,
	// 	tog.Basics4,
	// 	tog.Basics5,
	// 	tog.Basics6,
	// 	tog.Basics7,
	// 	tog.Basics8,
	// 	tog.Basics9,
	// 	tog.Basics10,
	// 	tog.Basics11,
	// 	tog.Basics12,
	// 	tog.Basics13,
	// 	tog.Basics14,
	// 	tog.Basics15,
	// 	tog.Basics16,
	// })

	// callFunctions("FlowControl", []func(){
	// 	tog.FlowControl1,
	// 	tog.FlowControl2,
	// 	tog.FlowControl3,
	// 	tog.FlowControl4,
	// 	tog.FlowControl5,
	// 	tog.FlowControl6,
	// 	tog.FlowControl7,
	// 	tog.FlowControl8,
	// 	tog.FlowControl9,
	// 	tog.FlowControl10,
	// 	tog.FlowControl11,
	// 	tog.FlowControl12,
	// 	tog.FlowControl13,
	// })

	callFunctions("Moretypes", []func(){
		tog.Moretypes1,
		tog.Moretypes2,
		tog.Moretypes3,
		tog.Moretypes4,
		tog.Moretypes5,
		tog.Moretypes6,
		tog.Moretypes7,
		tog.Moretypes8,
		tog.Moretypes9,
		tog.Moretypes10,
		tog.Moretypes11,
		tog.Moretypes12,
		tog.Moretypes13,
		tog.Moretypes14,
	})
}

func callFunctions(title string, f []func()) {
	for i := 0; i < len(f); i++ {
		fmt.Printf("\n---------- %s %d ----------\n\n", title, i+1)
		f[i]()
	}
}
