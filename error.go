package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Go programs express error state with error values. The error type is a built-in interface similar to the fmt.Stringer package.

// As with fmt.Stringer, the fmt package looks for the error interface when printing values.
func (m *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", m.When, m.What)
}

func run() error {
	return &MyError{ //
		time.Now(), // Setting the values in the struct
		"it didn't work",
	}
}
