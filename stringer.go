package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Stringer is implemented by any value that has a String method, which defines de "native" format for thar value
// The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age) // Inside the function, you associate the fields and customize how to print.
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	// When executed, as defined in the String() function, will show
	// Arthur Dent (42 years)
	// Zaphod Beeblebrox (9001 years)
}
