package main

import (
	"fmt"
	"math"
)

//	Functions san be used as values too. They can be passed around just like other values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) // In this case
}

func main() {

	// A variable with a function as a value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	// Using the hypot function as an argument in the compute function, since both of them have the same parameters
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}
