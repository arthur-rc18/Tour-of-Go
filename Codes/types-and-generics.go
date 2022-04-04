package main

import "fmt"

// Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets,
// before the function's arguments.

// This declaration means that 's' is a slice of any type that fulfills the built-in constraint 'comparable'.
func Index[T comparable](s []T, x T) int { // 'comparable' is a useful constraint that makes it possible to use the '==' and '!=' on values of the type.
	for i, v := range s {
		// v and x are type T, which has the comparable constraint, so we can use '==' here
		if v == x {
			return i
		}
	}
	return -1
}

// Using a genery type with the keyword 'any'
type List[T any] struct { //any is an alias for interface{} and is equivalent to interface{} in all ways.
	next *List[T]
	val  T
}

func main() {

	// T will assume a int value
	ss := []int{10, 19, 15, 200, 2001}
	fmt.Println(Index(ss, 200))

	// T will assume a string value
	si := []string{"test", "hello", "array", "generics"}
	fmt.Println(Index(si, "hello"))

	var x List[int]
	v := List[int]{next: &x, val: 10}

	fmt.Println(v)
	fmt.Println(x)

}
