package main

import "fmt"

type Test struct {
	X int //Inside of the struct you define the fields
	Y float32
	Z string
}

func ex() {
	fmt.Println(Test{10, 5.5, "Hello"}) //This will set the values of the fields in the struct
}

func instance() {

	v := Test{} //That's an easier way to set the fields using dot
	v.X = 15
	v.Y = 57.43
	v.Z = "Teting"

	fmt.Println(v.X, v.Y, v.Z)
}
