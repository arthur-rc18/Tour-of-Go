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

type V struct {
	S string
}

func (v *V) M() {
	if v == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(v.S)
}

type itc interface {
	M()
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
	// Calling the function with the interface already settted
	describe(a2)
	a2.tg()

	// If the concrete  value inside the interface itself is nil, the method will be called with a nil receiver.
	var i itc
	var v2 *V
	i = v2
	describe2(i)
	i.M() // Both functions will print <nil>, because there is no concrete value in the interface i.

	// Is it possible to create a empty interface
	var y interface{}
	describe3(y) // An empty interface may hold values of any type.

	y = 42 // Setting an int value
	describe3(y)

	y = "Testing" // Setting a string
	describe3(y)

	// Type assertions in interfaces
	// A type assertion provides access to an interface value's underlying concrete value

	var tAssert interface{} = "Hello" // Setting a string in the interface variable
	tai := tAssert.(string)           // Passing the interface's string as a value to tai
	fmt.Println(tai)

	// To test whether an interface value holds a specific type, a type assertation can return two values, the underlying value
	// and a boolean value that reports whether the assertion succeed

	// It's very simillar to verify the values in a map
	ve, ok := tAssert.(string) // If the interface tAssert holds a string, then 've' will be the value of the type and 'ok' will be true.
	fmt.Println(ve, ok)

	ve2, ok := tAssert.(float64) // If the interface does not hold the type, then 've2' will be the zero value of the type and ok will be false.
	fmt.Println(ve2, ok)         // In that case will return '0' and 'false', since the interface doesn't hold a float64 value
}

// Passing a interface as a value
// This function will only print the value and the Go-syntax representation of the type of the value that is setted in the interface.
func describe(i it) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i itc) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
