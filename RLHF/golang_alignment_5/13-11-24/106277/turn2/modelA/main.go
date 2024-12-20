package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Customer struct
type Customer struct {
	ID       int
	Age      int
	Gender   string
	Location string
}

// SegmentationResult struct
type SegmentationResult struct {
	CustomerID int
	Segment    string
}

// Customer struct remains the same

// SegmentCustomer function with error handling
func SegmentCustomer(ctx context.Context, customer Customer, results chan<- SegmentationResult) error {
	defer close(results)

	var segment string
	if customer.Age < 18 {
		segment = "Young"
	} else if customer.Age < 30 {
		segment = "Adult"
	} else {
		segment = "Senior"
	}

	select {
	case results <- SegmentationResult{CustomerID: customer.ID, Segment: segment}:
	default:
		return fmt.Errorf("result channel full")
	}
	return nil
}

func SegmentCustomersConcurrently(customers []Customer) []SegmentationResult {
	const numWorkers = 4

	var results []SegmentationResult
	wg := &sync.WaitGroup{}

	resultsChan := make(chan SegmentationResult, len(customers))

	// Create a context with a timeout to handle slow customers
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Worker goroutines with error handling
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case customer, ok := <-customers:
					if !ok {
						return
					}
					if err := SegmentCustomer(ctx, customer, resultsChan); err != nil {
						fmt.Println("error segmenting customer:", err)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// Send customers to worker goroutines
	go func() {
		for _, customer := range customers {
			customers <- customer
		}
		close(customers)
	}()

	wg.Wait()

	close(resultsChan)

	// Collect results
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func main() {
	// Sample customers
	customers := []Customer{
		// Add many more customers here
	}

	start := time.Now()
	results := SegmentCustomersConcurrently(customers)
	end := time.Now()

	fmt.Println("Segmentation Results:")
	for _, result := range results {
		fmt.Printf("Customer ID: %d, Segment: %s\n", result.CustomerID, result.Segment)
	}
	fmt.Printf("Execution Time: %s\n", end.Sub(start))
}
