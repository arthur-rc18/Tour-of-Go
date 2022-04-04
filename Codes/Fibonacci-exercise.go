package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

var x int = 1

func fibonacci() func() int {
	s1, s2 := 0, 0
	return func() int {
		s2 = x + s1
		s1 = x
		x = s2
		return s2

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
