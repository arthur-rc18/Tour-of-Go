package main

import (
	"fmt"
	"math"
)

type Vert struct {
	X, Y float64
}

func (v Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vert{3, 4}

	fmt.Println(v.Abs())

	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs2())

	f2 := Str{3, 4} //Making the instance of the struct
	f2.Scale(10)    // This function will be executed first when called
	fmt.Println(f2.Abs2())
}

// You can declare a method on non-struct types, too.
type Myfloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method.
func (f Myfloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Str struct {
	X, Y float64
}

// This is the original function
func (s Str) Abs2() float64 {
	return math.Sqrt(s.X*s.X + s.Y*s.Y)
}

// 	Methods can be declared with pointers receivers. This means the receiver type has the literal syntax *T for some type T.
// The pointer will make the conection to the struct and its fields, modifying its values.
func (s *Str) Scale(f float64) {
	s.X = s.X * f
	s.Y = s.Y * f
}
