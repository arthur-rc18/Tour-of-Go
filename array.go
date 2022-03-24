package main

import "fmt"

func main() {

	var a [2]string // Example of creating an array and setting the values
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // How to print an array in two differente manners
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // You can set the values of an array like that
	fmt.Println(primes)
	i := [2]int{5, 9} // Example of an array in a function
	fmt.Println(sum(i[0], i[1]))

	var j []int = primes[2:4] // This is called "slice", a slice is formed by specifyin two indices, a low and a high bound, separated by a colon.
	fmt.Println(j)            // This selects a half-open range which includes the first element, but excludes the last one.

	names := [4]string{"John", "Paul", "George", "Ringo"}
	x := names[0:2] // x will have "John" and "Paul"
	y := names[1:3] // y will have "Paul" and "George"
	fmt.Println(x, y)

	y[0] = "XXX"       // Since y have only two values, "Paul" is in the 0 position and "George" in the 1, so only "Paul" is going to change
	fmt.Println(names) // When the array names is printed will show [John XXX George Ringo]

	s := []struct { // Example of a literal slice
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true}, {7, true},
	}
	fmt.Println(s)

}

func sum(x, y int) int {
	return x + y
}
