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

func (v *ver) tg() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f Myfl) tg() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type it interface {
	tg() float64
}

func main() {
	var a it
	f := Myfl(math.Sqrt2)
	v := ver{3, 4}

	a = f
	a = &v

	fmt.Println(a.tg())
}
