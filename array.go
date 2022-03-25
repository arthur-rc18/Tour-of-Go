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

	s1 := []int{2, 3, 6, 9, 6, 20, 21}
	s1 = s1[:5] // When there is no value in the left side, the position of the right side will "cute" the array and take all values at the right

	s2 := []int{69, 5, 23, 9, 3, 9, 4}
	size(s2) // Calling the size function down at the code

	//Slice the slice to give it zero length
	s2 = s2[:0]
	size(s2)

	// Extend its length
	s2 = s2[:4]
	size(s2)

	// Drop its first two values
	s2 = s2[2:]
	size(s2)

	// The 'make' function allocates and initializes an object of type slice, map or chan
	ar := make([]int, 5) // Here you specify the type of the array and then its lenght
	size(ar)

	ar2 := make([]int, 0, 10) // Sometimes the function can have 3 arguments, the third argument it's the capacity
	size(ar2)

}

func sum(x, y int) int {
	return x + y
}

func size(s []int) {
	fmt.Printf("The size of the array is %d, the capacity is %d and the slice is %v\n", len(s), cap(s), s) // This function will return the lenght, the capacity
	// of an array and the array or the slice of it
}
