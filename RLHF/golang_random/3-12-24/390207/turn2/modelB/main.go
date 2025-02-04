package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	incrementCounter := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Increment()
		}
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrementCounter()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
