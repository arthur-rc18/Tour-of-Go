package main

import (
	"fmt"
	"time"
)

// Channels are a typed conduit through which you can send and recieve values with the channel operator '<-'

func plus(s []int, c chan int) { // The key word 'chan' indicates that the variable 'c' it's a channel
	sum := 0
	for _, i := range s {
		sum += i
	}

	c <- sum // The arrow indicates the direction of the data
}

func main() {

	s := []int{2, 7, 9, 11, -9, 5}

	// Like maps and slices, channels must be created before use
	c := make(chan int) // The buffer is initialized with the specified buffer capacity. If zero, or the size is omitted, the channel is unbuffered

	go plus(s[:len(s)/2], c) // c <- 18
	time.Sleep(1000 * time.Millisecond)
	go plus(s[len(s)/2:], c) // c <- 7

	x, y := <-c, <-c // 'x' will receive the value of the first c, and 'y' will receive the value of the second c

	fmt.Println(x, y, x+y)
}
