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

	go plus(s[:len(s)/2], c)            // c <- 18
	time.Sleep(1000 * time.Millisecond) // Without this time, 'x' would receive the value of the second c and 'y' would receice the first c
	go plus(s[len(s)/2:], c)            // c <- 7

	x, y := <-c, <-c // 'x' will receive the value of the first c, and 'y' will receive the value of the second c

	fmt.Println(x, y, x+y)

	// Buffered Channels
	// Channels can be buffered. Provide the buffer length as the second argument to 'make' to initialize a buffered channel

	ch := make(chan int, 3)

	ch <- 5
	ch <- 9
	ch <- 75 // The channel has to provide enough values based on the length
	// In that it was three values

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	c2 := make(chan int, 10)
	// Calling the fib function created previously
	go fib(cap(c2), c2)
	for i := range c2 { // The loop ' for i := range c2" receives values from the channel repeatedly until is closed
		fmt.Println(i)
	}

	// Instancing a function inside the main, to show how select works
	// This part of the code will return a panic, this will be solved and talked in other code
	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib2(c3, quit)

}

// Demonstration of how the 'close' function works

func fib(n int, c chan int) { // This function returns the n's number of Fibonacci
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Select in channels
// Selects acts like a switch case. The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		time.Sleep(2 * time.Second)
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
