package main

import (
	"fmt"
	"sync"
)

// SharedResource represents a resource accessed by multiple goroutines
type SharedResource struct {
	mu      sync.Mutex
	counter int
}

func (r *SharedResource) Increment() {
	r.mu.Lock()
	r.counter++
	r.mu.Unlock()
}

func main() {
	resource := &SharedResource{}

	// Goroutines simulating concurrent work
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				resource.Increment()
			}
		}()
	}

	// Wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10)
	wg.Wait()

	fmt.Println("Final Counter Value:", resource.counter) // Expected output: 10000
}
