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
	return c.data[key], c.data[key] != ""
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
	var wg sync.WaitGroup
	wg.Add(10) // Add one wait group per 100 keys

	for key := 0; key < 1000; key++ {
		go func(key int) {
			// Check in L1 cache first
			if value, found := l1Cache.get(key); found {
				fmt.Printf("Key %d found in L1 cache: %s\n", key, value)
				return
			}

			// Check in L2 cache
			if value, found := l2Cache.get(key); found {
				fmt.Printf("Key %d found in L2 cache: %s\n", key, value)
				updateCacheLevels(key, value, l1Cache)
				return
			}

			// Check in L3 cache
			if value, found := l3Cache.get(key); found {
				fmt.Printf("Key %d found in L3 cache: %s\n", key, value)
				updateCacheLevels(key, value, l2Cache, l1Cache)
				return
			}

			// Fetch from external source if not found in all caches
			value := fetchFromExternal(key)
			fmt.Printf("Key %d fetched from external source: %s\n", key, value)
			updateCacheLevels(key, value, l3Cache, l2Cache, l1Cache)

			wg.Done()
		}(key)

		if key%100 == 0 {
			wg.Wait() // Wait for the 100 keys to complete before moving on
			wg.Add(10)
		}
	}

	wg.Wait() // Wait for any remaining goroutines to finish
}
