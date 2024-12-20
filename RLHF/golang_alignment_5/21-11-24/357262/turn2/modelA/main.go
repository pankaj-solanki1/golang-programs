package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel
	ch := make(chan int)

	go func() {
		fmt.Println("Goroutine 1 starting...")
		ch <- 1 // Sender blocks until receiver is ready
		fmt.Println("Goroutine 1 finished sending.")
	}()

	go func() {
		fmt.Println("Goroutine 2 starting...")
		msg := <-ch // Receiver blocks until sender sends a message
		fmt.Println("Goroutine 2 received:", msg)
		fmt.Println("Goroutine 2 finished.")
	}()

	fmt.Println("Main goroutine starting...")
	time.Sleep(time.Second)
	fmt.Println("Main goroutine finished.")
}
