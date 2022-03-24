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

}

func sum(x, y int) int {
	return x + y
}
