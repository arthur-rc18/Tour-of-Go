package main

import "fmt"

// Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets,
// before the function's arguments.

// This declaration means that 's' is a slice of any type that fulfills the built-in constraint 'comparable'.
func Index[T comparable](s []T, x T) int { // 'comparable' is a useful constraint that makes it possible to use the '==' and '!=' on values of the type.
	for i, v := range s {

		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct { // any is an alias for interface{} and is equivalent to interface{} in all ways.
	next *List[T]
	val  T
}

func main() {

	ss := []int{10, 19, 15, 200, 2001}
	fmt.Println(Index(ss, 200))

	si := []string{"test", "hello", "array", "generics"}
	fmt.Println(Index(si, "hello"))

	var x List[int]
	v := List[int]{next: &x, val: 105}
	fmt.Println(v)
}
