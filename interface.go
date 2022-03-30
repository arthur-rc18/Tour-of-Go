package main

import (
	"fmt"
	"math"
)

// Initializing the type
type Myfl float64

// Creating the struct
type ver struct {
	X, Y float64
}

type ver2 struct {
	X, Y float64
}

// Creating the functions with the same name but different parameters
func (v *ver) tg() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *ver2) tg() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f Myfl) tg() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type it interface {
	tg() float64 // The parameter of the interface is the function already created.
	// The function that is common in both the struct and the Myfl
}

func main() {
	var a it // Here it's creating the instance of the interface
	// Interface can be implemented in a implicity way

	f := Myfl(math.Sqrt2)
	v := ver{3, 4} // Setting the types already created

	a = f
	a = &v
	fmt.Println(a.tg())

	var t *ver2
	var a2 it
	a2 = t

	a2 = &ver2{3, 4}
	describe(a2)
	a2.tg()

	// Calling the function with the interface already settted

}

// Passing a interface as a value
// This function will only print the value and the Go-syntax representation of the type of the value that is setted in the interface.
func describe(i it) {
	fmt.Printf("(%v, %T)", i, i)
}
