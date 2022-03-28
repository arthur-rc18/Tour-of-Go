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
}
