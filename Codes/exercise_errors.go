package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	z := float64(2.5)
	for i := 0; i < 1; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	var g float64
	x, _ := Sqrt(g)
	if x < 0 {
		return ErrNegativeSqrt(-2).Error()

	} else {
		Sqrt(x)
	}
	return ""
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-9))
}
