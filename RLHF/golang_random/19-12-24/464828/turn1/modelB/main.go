
package main  
import (  
    "fmt"
    "time"
)

// A function that returns multiple values
func multipleReturnValues() (int, int, string) {  
    return 10, 20, "Hello"
}

// A function that accepts multiple values as parameters
func acceptMultipleValues(a int, b int, c string) {  
    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
}

func main() {  
    start := time.Now()

    // Calling the function that returns multiple values and storing the results in individual variables
    a, b, c := multipleReturnValues()
    acceptMultipleValues(a, b, c)

    // Calling the function that returns multiple values and storing the results in a slice
    result := multipleReturnValues()
    acceptMultipleValues(result[0], result[1], result[2].(string))

    fmt.Println("Time taken using multiple return values: ", time.Since(start))

    start = time.Now()

    // Using named return values to improve code readability and reduce the risk of errors
    func namedReturnValues() (x int, y int, z string) {  
        x = 10
        y = 20
        z = "Hello"
        return
    }
    a, b, c = namedReturnValues()
    acceptMultipleValues(a, b, c)

    fmt.Println("Time taken using named return values: ", time.Since(start))
}  
