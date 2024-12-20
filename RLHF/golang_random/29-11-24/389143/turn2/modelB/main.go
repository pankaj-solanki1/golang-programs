package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/29-11-24/389143/turn2/modelB/calculations"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define our routes
	v1 := router.Group("/v1")
	{
		v1.GET("/calculate/:operator/:a/:b", calculations.CalculateHandler)
	}

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
