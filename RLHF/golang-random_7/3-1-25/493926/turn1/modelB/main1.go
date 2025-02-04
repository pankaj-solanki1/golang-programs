package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Mock database interface
type MockDB interface {
	GetData() (string, error)
}

// Mock database implementation
type MockDBImpl struct{}

func (m *MockDBImpl) GetData() (string, error) {
	time.Sleep(time.Second) // Simulate some I/O delay
	return "Sample Data", nil
}

// Cloud-native application function
func main() {
	db := &MockDBImpl{}

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Number of concurrent read and write operations
	numOperations := 10

	// Start reading data from the database concurrently
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			data, err := db.GetData()
			if err != nil {
				log.Printf("Error reading data: %v", err)
				return
			}
			fmt.Println("Read data:", data)
		}()
	}

	// Start writing data to a file concurrently
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			// Simulate writing data to a file
			time.Sleep(time.Second)
			fmt.Println("Wrote data to file.")
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}
