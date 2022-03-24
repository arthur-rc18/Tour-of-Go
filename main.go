package main

import (
	"fmt"
	"runtime"

	"example.com/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Opt())
	//fmt.Println(runtime.GOOS)
	x := runtime.GOOS
	var c = &x
	fmt.Println(*c)
	*c = "windows"
	fmt.Println(x)
}
