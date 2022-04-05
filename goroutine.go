package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i <= 4; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(430 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}
