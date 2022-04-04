package main

import (
	"fmt"
	"io" // The io package specifies the io.Reader interface, which represents the read of a stream of data.
	"strings"

	"golang.org/x/tour/reader"
)

func main() {
	// NewReader returns a new Reader reading from s.
	v := strings.NewReader("Hello, Reader!")

	// The make function
	b := make([]byte, 8)
	for {
		// Read populates the given byte slice with data and return the number of bytes populated and an error value. It returns an io.EOF error when
		// the stream ends.
		n, err := v.Read(b) // Read implements the io.Reader interface
		fmt.Printf("n = %v err = %v b = %v\n ", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// EOF = End Of File
		if err == io.EOF {
			break
		}
		reader.Validate(v)
	}
}
