package main

import (
	"fmt"
	"sync"
	"time"
)

// task represents a task to be processed
type task struct {
	id   int
	data string
}

// worker is a function that processes tasks from a channel
type worker func(chan task)

// startWorkers starts a specified number of workers to process tasks from a channel
func startWorkers(wg *sync.WaitGroup, numWorkers int, tasks chan task, worker worker) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(tasks)
		}()
	}
}

// processTask processes a task with a simulated delay
func processTask(tasks chan task) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d: %s\n", task.id%10, task.id, task.data)
		// Simulate task processing time
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Create a channel to hold tasks
	tasks := make(chan task)

	var wg sync.WaitGroup
	numWorkers := 5
	wg.Add(1)

	// Start workers to consume tasks from the channel
	go func() {
		defer wg.Done()
		startWorkers(&wg, numWorkers, tasks, processTask)
	}()

	// Generate and send tasks to the channel
	for i := 1; i <= 100; i++ {
		tasks <- task{id: i, data: fmt.Sprintf("Task %d", i)}
	}

	// Close the task channel to indicate the end of tasks
	close(tasks)

	wg.Wait()
	fmt.Println("All tasks completed.")
}