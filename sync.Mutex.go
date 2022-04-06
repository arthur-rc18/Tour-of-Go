package main

import (
	"fmt"
	"sync" // This is library respomsible for this
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// When communication isn't needed and we just want to make sure only one goroutine can access a variable at a time to avoid conflicts,
// this is called mutual exclusion

func (c *SafeCounter) Inc(key string) {

	c.mu.Lock()
	//Lock so only one goroutine at a time can access the c.v

	c.v[key]++
	c.mu.Unlock() // Unlocking the function
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()

	defer c.mu.Unlock()
	return c.v[key]
}

func main() {

	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("Somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
