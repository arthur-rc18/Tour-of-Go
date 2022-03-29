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
	fmt.Println(f.Abs2()-)

}

// You can declare a method on non-struct types, too.
type Myfloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method.
func (f Myfloat) Abs2() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
		
}	


