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

	fmt.Println(mp) // When executed, will print the keys to its values
	// map[apple:5 grape:9 orange:2]

	// map works with others functions
	fmt.Println(len(mp)) // This will show the size of the map

	// To add another value inside the map
	mp["pear"] = 900

	// To delete a value
	delete(mp, "apple")

	fmt.Println(mp)
}
