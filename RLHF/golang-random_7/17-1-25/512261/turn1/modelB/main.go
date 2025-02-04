package main

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// A sample third-party library function that might not handle defer correctly.
func thirdPartyFunction() error {
	defer fmt.Println("Defer from thirdPartyFunction.")
	return errors.New("something went wrong")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// Call the third-party function
		err := thirdPartyFunction()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	wg.Wait()
	fmt.Println("Main function finished.")
}
