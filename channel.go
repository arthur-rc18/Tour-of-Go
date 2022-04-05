package main

import (
	"fmt"
	"time"
)

func plus(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}

	c <- sum
}

func main() {

	s := []int{2, 7, 9, 11, -9, 5}

	c := make(chan int)

	go plus(s[:len(s)/2], c)
	time.Sleep(1000 * time.Millisecond)
	go plus(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
