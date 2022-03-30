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

// Creating the functions with the same name but different parameters
func (v *ver) tg() float64 {
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
	var a2 it = &ver{4, 3}
	f := Myfl(math.Sqrt2)
	v := ver{3, 4} // Setting the types already created

	a = f
	a = &v

	fmt.Println(a.tg())
	fmt.Println(a2.tg())
}
