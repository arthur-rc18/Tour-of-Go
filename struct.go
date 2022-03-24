package main

import "fmt"

type Struct struct {
	X int //Inside of the struct you define the fields
	Y float32
	Z string
}

func execute() {
	fmt.Println(Struct{10, 5.5, "Hello"}) //This will set the values of the fields in the struct
}

func setting() {

	v := Struct{} //That's an easier way to set the fields using dot
	v.X = 15
	v.Y = 57.43
	v.Z = "Teting"

	fmt.Println(v.X, v.Y, v.Z)
}

var (
	v1 = Struct{1, 5.5, "Testing"} // has type Struct
	v2 = Struct{X: 5}              // Y:0 and Z:"" is implicit
	v3 = Struct{}                  // X:0, Y:0 and Z:""
	v4 = &Struct{10, 5, "Hi"}      // Has type *Struct
)
