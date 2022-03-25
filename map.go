package main

import "fmt"

func main() {

	// A map maps keys to values
	// In this example, it's mapping a string to int values
	var mp map[string]int = map[string]int{
		"apple":  5,
		"grape":  9,
		"orange": 2,
	}

	fmt.Println(mp)
}
