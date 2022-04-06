package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond) // If you don't specify the time, it will be instantly.
		fmt.Printf("%d ", i)
	}
}

func alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(430 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

// A goroutine is a lightweight thread managed by the Go runtime.
func main() {
	go numbers()                        // This is how you start a goroutine
	go alphabet()                       // These two goroutine will be executed at the same time, but with a difference of 30 Milliseconds, as setted before.
	time.Sleep(3000 * time.Millisecond) // This means that the last command will take 3 seconds to be executed.
	fmt.Println("Main terminated")      // The goroutine will be terminated after the main goroutine, even if has or not completed the execution.
}
