package main

import (
	"fmt"
	"sync"
	"time"
)

type cache struct {
	data map[int]string
	m    sync.RWMutex
}

func newCache() *cache {
	return &cache{
		data: make(map[int]string),
	}
}

func (c *cache) get(key int) (string, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	value, found := c.data[key]
	return value, found
}

func (c *cache) set(key int, value string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[key] = value
}

func fetchFromExternal(key int) string {
	// Simulate external data fetching by adding a delay
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("External data for key %d", key)
}

func updateCacheLevels(key int, value string, caches ...*cache) {
	for _, c := range caches {
		c.set(key, value)
	}
}

func main() {
	// Initialize caches
	l1Cache := newCache()
	l2Cache := newCache()
	l3Cache := newCache()

	// Simulated data range
	numGoroutines := 10
	dataRange := 10000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for key := goroutineID * dataRange; key < (goroutineID+1)*dataRange; key++ {
				// Check in L1 cache first
				if value, found := l1Cache.get(key); found {
					continue
				}

				// Rest of the caching logic remains the same
				// ...
				// ...
			}
		}(i)
	}

	wg.Wait()
	end := time.Now()
	fmt.Println("Total time taken:", end.Sub(start))
}
